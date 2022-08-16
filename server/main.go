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

	httpServer, err := NewServer(ctx)
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

// implements mux.HandlerFunc, we ignore path params in this scenario.
func handler(h http.Handler) func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		h.ServeHTTP(w, r)
	}
}

// implements mux.HandlerFunc, we ignore path params in this scenario.
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

type HTTPServer struct {
	server *http.Server
}

// NewServer returns an HTTPServer instance with a handler attached.
// Note, handlers should implement a timeout to avoid running into transport
// layer timeouts.
// Add HTTP middleware here.
func NewServer(ctx context.Context) (*HTTPServer, error) {
	mux := runtime.NewServeMux()

	// openapi documentation handler
	fs := http.FileServer(http.Dir("./apis/ui"))
	if err := mux.HandlePath("GET", openApiPrefix+".json", handleOpenapiDescription); err != nil {
		return nil, fmt.Errorf("could not register openapi json endpoint: %v", err)
	}
	if err := mux.HandlePath("GET", openApiPrefix+"/*", handler(http.StripPrefix(openApiPrefix, fs))); err != nil {
		return nil, fmt.Errorf("could not register openapi endpoint: %v", err)
	}

	// regiester with gRPC endpoint
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterMatchServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		return nil, fmt.Errorf("could not register mux router with gRPC endpoint: %v", err)
	}

	// add config
	port := "8081"

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      mux,
		ReadTimeout:  30 * time.Second, // deadline for reading request body
		WriteTimeout: 30 * time.Second, // deadline for ServeHTTP
	}
	return &HTTPServer{server: server}, nil

}

func (s *HTTPServer) Run() error {
	return s.server.ListenAndServe()
}

// Shutdown stops accepting new requests and waits for the running
// ones to finish before returning. See net/http docs for details.
// The provided context should have a timeout.
func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
