package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	ent "github.com/fbngrm/around-home/ent"
	apiv1 "github.com/fbngrm/around-home/gen/proto/go/match/v1"
	gw "github.com/fbngrm/around-home/gen/proto/go/match/v1"
	"github.com/fbngrm/around-home/pkg/database"
	"github.com/fbngrm/around-home/pkg/location"
	"github.com/fbngrm/around-home/pkg/match"
	"github.com/fbngrm/around-home/pkg/partner"
	"github.com/fbngrm/around-home/server/api"
	"github.com/fbngrm/around-home/server/internal/postgres"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const openApiPrefix = "/openapi"

var (
	// command-line options:
	// gRPC server endpoint
	grpcEndpoint = flag.String("grpc-server-endpoint", "localhost:8080", "gRPC server endpoint")
)

func main() {
	// flag.StringVar(&grpcEndpoint, "grpc-endpoint", "localhost:8080", "gRPC server endpoint")
	flag.Parse()

	// TODO: use functional options here to avoid empty config
	db, err := postgres.Connect(database.Config{})
	if err != nil {
		log.Printf("could not connect database %v\n", err)
		os.Exit(1)
	}

	// 	ctx, cancel := context.WithCancel(context.Background())

	// quitCh := make(chan os.Signal, 1)
	// // interrupt signal sent from terminal
	// signal.Notify(quitCh, os.Interrupt)
	// // sigterm signal sent from kubernetes or docker
	// signal.Notify(quitCh, syscall.SIGTERM)

	// go func() {
	// 	sig := <-quitCh
	// 	log.Printf("received interrupt %s, shutting down...\n", sig)
	// 	cancel()
	// }()

	go func() {
		for i := 0; i < 10; i++ {
			if err := runHTTP(); err != nil {
				log.Printf("%+v", err)
			}
			time.Sleep(time.Millisecond * 100)
		}
		log.Fatal("HTTP Gateway shutdown")
	}()

	if err := runGRPC(db); err != nil {
		log.Fatal(err)
	}
	log.Fatal("GRPC shutdown")
}

func runHTTP() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()

	// openapi documentation.
	fs := http.FileServer(http.Dir("./apis/ui"))
	if err := mux.HandlePath("GET", openApiPrefix+".json", handleOpenapiDescription); err != nil {
		log.Fatal("add to mux openapi json", err)
	}
	if err := mux.HandlePath("GET", openApiPrefix+"/*", handleHandler(http.StripPrefix(openApiPrefix, fs))); err != nil {
		log.Fatal("add to mux openapi", err)
	}

	fmt.Println(*grpcEndpoint)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterMatchServiceHandlerFromEndpoint(ctx, mux, *grpcEndpoint, opts)
	if err != nil {
		return err
	}
	listenOn := ":8081"
	log.Println("HTTP listening on", listenOn)
	return http.ListenAndServe(listenOn, mux)
}

func runGRPC(db *ent.Client) error {
	partnerRepo := &partner.Repo{DB: db}
	locationService := location.NewService(&location.HaversineCalculator{})
	matchService := match.NewMatchService(partnerRepo, locationService)
	api := api.NewApi(matchService)

	server := grpc.NewServer()
	apiv1.RegisterMatchServiceServer(server, api)

	listener, err := net.Listen("tcp", *grpcEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", *grpcEndpoint, err)
	}

	log.Println("gRPC listening on", *grpcEndpoint)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

func handleHandler(h http.Handler) func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		h.ServeHTTP(w, r)
	}
}

func handleOpenapiDescription(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	content, err := os.ReadFile("gen/openapiv2/match/v1/match.swagger.json")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	if _, err = w.Write(content); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
}
