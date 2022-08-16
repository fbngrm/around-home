package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/fbngrm/around-home/ent"
	apiv1 "github.com/fbngrm/around-home/gen/proto/go/match/v1"
	"github.com/fbngrm/around-home/pkg/location"
	"github.com/fbngrm/around-home/pkg/match"
	"github.com/fbngrm/around-home/pkg/partner"
	"github.com/fbngrm/around-home/server/api"
	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
}

// NewServer returns an Server instance with a handler attached.
// Note, handlers should implement a timeout to avoid running into transport
// layer timeouts.
func NewServer(ctx context.Context, db *ent.Client) (*Server, error) {
	// deps
	partnerRepo := &partner.Repo{DB: db}
	locationService := location.NewService(&location.HaversineCalculator{})
	matchService := match.NewMatchService(partnerRepo, locationService)
	api := api.NewApi(matchService)

	// note, not a production ready config
	server := grpc.NewServer()
	apiv1.RegisterMatchServiceServer(server, api)

	return &Server{
		server: server,
	}, nil
}

func (s *Server) Run(grpcEndpoint string) error {
	// TODO: add config
	listener, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", grpcEndpoint, err)
	}

	log.Println("gRPC listening on", grpcEndpoint)
	if err := s.server.Serve(listener); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}
	return nil
}

// stops the gRPC server gracefully. It stops the server from accepting new
// connections and RPCs and blocks until all the pending RPCs are  finished.
func (s *Server) Shutdown() {
	s.server.GracefulStop()
}
