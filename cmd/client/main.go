package main

import (
	"context"
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	apiv1 "github.com/fbngrm/around-home/gen/proto/go/match/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// connectTo := "127.0.0.1:8080"
	connectTo := "172.22.0.4:8080"
	conn, err := grpc.Dial(
		connectTo,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return fmt.Errorf("failed to connect to match service on %s: %w", connectTo, err)
	}
	log.Println("connected to", connectTo)

	svc := apiv1.NewMatchServiceClient(conn)
	ctx := context.Background()

	log.Println("fetching")
	matches, err := svc.MatchPartnersWithRequest(ctx, &apiv1.MatchPartnersWithRequestInput{
		Material: "wood",
		Location: "52.532566/13.396261",
	})
	if err != nil {
		return fmt.Errorf("could not get matches: %v\n", err)
	}
	spew.Dump(matches)

	return nil
}
