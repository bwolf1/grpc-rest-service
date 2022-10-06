package echo

import (
	"context"
	"time"

	pb "github.com/bwolf1/grpc-rest-kubernetes/proto"
)

type Server struct {
	pb.UnimplementedEchoerServer
}

func (s *Server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Echo:      in.GetWord(),
		Timestamp: time.Now().UTC().String(),
	}, nil
}
