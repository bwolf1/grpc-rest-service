package cmd

import (
	"context"
	"log"
	"time"

	pb "github.com/bwolf1/grpc-rest-kubernetes/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var word string

var grpcClientCmd = &cobra.Command{
	Use:     "grpcClient",
	Aliases: []string{"client"},
	Short:   "Run the gRPC client",
	Run: func(cmd *cobra.Command, args []string) {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.ReadInConfig()

		// Create the client connection to the server.
		conn, err := grpc.Dial(
			viper.GetString("grpc.environment.development.hostname")+":"+viper.GetString("grpc.environment.development.port"),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewEchoerClient(conn)

		// Create a context for the RPC (remote procedure call).
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		// Call the Echo RPC method on the server.
		r, err := c.Echo(ctx, &pb.EchoRequest{Word: word})
		if err != nil {
			log.Fatalf("unable to echo: %v", err)
		}

		// Log the RPC response.
		log.Printf("%v", r)
	},
}

func init() {
	grpcClientCmd.Flags().StringVarP(&word, "word", "w", "default word", "Word to be echoed")
	rootCmd.AddCommand(grpcClientCmd)
}
