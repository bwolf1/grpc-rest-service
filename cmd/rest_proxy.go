package cmd

import (
	"context"
	"log"
	"net/http"

	"github.com/bwolf1/grpc-rest-kubernetes/pkg/service/echo"
	pb "github.com/bwolf1/grpc-rest-kubernetes/proto"
	"github.com/davecgh/go-spew/spew"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var restProxyCmd = &cobra.Command{
	Use:     "restProxy",
	Aliases: []string{"client"},
	Short:   "Run the REST proxy",
	Run: func(cmd *cobra.Command, args []string) {
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.ReadInConfig()

		// Start the server.
		grpcServer := grpc.NewServer()
		pb.RegisterEchoerServer(grpcServer, &echo.Server{})
		conn, err := grpc.DialContext(
			context.Background(),
			"0.0.0.0:50051", // TODO (bwolf): Get the hostname and port from viper
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalln("Failed to dial server:", err)
		}
		router := runtime.NewServeMux()
		if err = pb.RegisterEchoerHandler(context.Background(), router, conn); err != nil {
			log.Fatalln("Failed to register gateway:", err)
		}
		http.ListenAndServe(
			":8080", // TODO (bwolf): Get the hostname and port from viper
			httpRouter(router),
		)
	},
}

func httpRouter(httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		spew.Println("Debug: Before ServeHTTP")
		httpHandler.ServeHTTP(w, r)
		spew.Println("Debug: After ServeHTTP")
	})
}

func init() {
	rootCmd.AddCommand(restProxyCmd)
}
