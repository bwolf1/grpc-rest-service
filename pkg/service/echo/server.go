package echo

import (
	"context"
	"log"
	"time"

	pb "github.com/bwolf1/grpc-rest-kubernetes/proto"
)

// Server implements the gRPC EchoerServer server.
type Server struct {
	pb.UnimplementedEchoerServer
}

// Echo takes a single word and returns it to the caller with a UTC timestamp.
func (s *Server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("echoing %s", in.Word)
	return &pb.EchoResponse{
		Echo:      in.GetWord(),
		Timestamp: time.Now().UTC().Truncate(1000 * time.Millisecond).String(),
	}, nil
}
