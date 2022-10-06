package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/bwolf1/grpc-rest-kubernetes/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// TODO move these values to the config after adding Viper.
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	word = flag.String("word", "", "Word to echo")
)

func main() {
	flag.Parse()
	// Create the client connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewEchoerClient(conn)

	// Create a context for the RPC (remote procedure call).
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call the Echo RPC method on the server.
	r, err := c.Echo(ctx, &pb.EchoRequest{Word: *word})
	if err != nil {
		log.Fatalf("unable to echo: %v", err)
	}

	// Log the RPC response.
	log.Printf("%v", r)
}
