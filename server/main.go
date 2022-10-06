package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/bwolf1/grpc-rest-kubernetes/proto"
	"google.golang.org/grpc"

	"github.com/bwolf1/grpc-rest-kubernetes/pkg/service/echo"
)

var (
	// TODO move these values to the config after adding Viper.
	port = flag.Int("port", 50051, "The server port")
)

// type server struct {
// 	pb.UnimplementedEchoerServer
// }

// TODO move this out of the main pkg and make a test for it after adding cobra.
// Echo takes a single word and returns it to the caller with a UTC timestamp.
// func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
// 	return &pb.EchoResponse{
// 		Echo:      in.GetWord(),
// 		Timestamp: time.Now().UTC().String(),
// 	}, nil
// }

func main() {
	// Start the server.
	flag.Parse()
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoerServer(s, &echo.Server{})
	log.Printf("server listening at %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("server failure: %v", err)
	}
}
