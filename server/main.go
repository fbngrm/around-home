package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	ent "github.com/fbngrm/around-home/ent"
	apiv1 "github.com/fbngrm/around-home/gen/proto/go/match/v1"
	"github.com/fbngrm/around-home/pkg/database"
	"github.com/fbngrm/around-home/pkg/location"
	"github.com/fbngrm/around-home/pkg/match"
	"github.com/fbngrm/around-home/pkg/partner"
	"github.com/fbngrm/around-home/server/api"
	"github.com/fbngrm/around-home/server/http"
	"github.com/fbngrm/around-home/server/internal/postgres"

	"google.golang.org/grpc"
)

var (
	grpcEndpoint string
)

func main() {
	flag.StringVar(&grpcEndpoint, "grpc-endpoint", "localhost:8080", "gRPC server endpoint")
	flag.Parse()

	// TODO: use functional options here to avoid empty config
	db, err := postgres.Connect(database.Config{})
	if err != nil {
		log.Printf("could not connect database %v\n", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())

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

	httpServer, err := http.NewServer(ctx, grpcEndpoint)
	if err != nil {
		log.Printf("could not init http server: %v\n", err)
		os.Exit(1)
	}

	go func() {
		for i := 0; i < 10; i++ {
			if err := httpServer.Run(); err != nil {
				log.Printf("%+v", err)
			}
			time.Sleep(time.Millisecond * 100)
		}
		log.Fatal("HTTP Gateway shutdown")
	}()

	if err := runGRPC(db); err != nil {
		log.Fatal(err)
	}

	cancel()
	log.Fatal("GRPC shutdown")
}

func runGRPC(db *ent.Client) error {
	partnerRepo := &partner.Repo{DB: db}
	locationService := location.NewService(&location.HaversineCalculator{})
	matchService := match.NewMatchService(partnerRepo, locationService)
	api := api.NewApi(matchService)

	// note, not a production ready config
	server := grpc.NewServer()
	apiv1.RegisterMatchServiceServer(server, api)

	listener, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", grpcEndpoint, err)
	}

	log.Println("gRPC listening on", grpcEndpoint)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	return nil
}

// func runHTTP() error {
// 	ctx := context.Background()
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	mux := runtime.NewServeMux()

// 	// openapi documentation.
// 	fs := http.FileServer(http.Dir("./apis/ui"))
// 	if err := mux.HandlePath("GET", openApiPrefix+".json", handleOpenapiDescription); err != nil {
// 		log.Fatal("add to mux openapi json", err)
// 	}
// 	if err := mux.HandlePath("GET", openApiPrefix+"/*", handler(http.StripPrefix(openApiPrefix, fs))); err != nil {
// 		log.Fatal("add to mux openapi", err)
// 	}

// 	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
// 	err := gw.RegisterMatchServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
// 	if err != nil {
// 		return err
// 	}
// 	listenOn := ":8081"
// 	log.Println("HTTP listening on", listenOn)
// 	return http.ListenAndServe(listenOn, mux)
// }
