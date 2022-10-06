package cmd

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/bwolf1/grpc-rest-kubernetes/pkg/service/echo"
	pb "github.com/bwolf1/grpc-rest-kubernetes/proto"
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
			viper.GetString(
				"grpc.environment.development.hostname")+":"+viper.GetString("grpc.environment.development.port"),
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
			":"+viper.GetString("rest.environment.development.port"),
			httpGrpcRouter(grpcServer, router),
		)
	},
}

func init() {
	rootCmd.AddCommand(restProxyCmd)
}

func httpGrpcRouter(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})
}
