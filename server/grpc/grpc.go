package grpc

import (
	"context"
	"fmt"
	"net"

	pbhighscore "github.com/maaram93/client/proto/v1"
	"google.golang.org/grpc"
)

// Grpc defines the server
type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 9999999999.0

// NewServer creates a new server with input address.
func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

// SetHighScore compares the input and updates  high_score.
func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighscoreRequest) (*pbhighscore.SetHighscoreResponse, error) {
	fmt.Println("SetHighScore endpoint is called")
	HighScore = input.HighScore
	return &pbhighscore.SetHighscoreResponse{
		Flag: true,
	}, nil
}

// GetHighScore returns high_score.
func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighscoreRequest) (*pbhighscore.GetHighscoreResponse, error) {
	fmt.Println("GetHighScore endpoint is called")
	return &pbhighscore.GetHighscoreResponse{
		HighScore: HighScore,
	}, nil
}

// ListenAndServe - listens and server at a paticular port
func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		fmt.Println("failed to open tcp port", err)
	}

	serverOptns := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOptns...)

	pbhighscore.RegisterGameServer(g.srv, g)

	fmt.Println("starting grpc server for highscore microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		fmt.Println("fialed to start grpc server for hignscore microservice")
	}

	return nil
}
