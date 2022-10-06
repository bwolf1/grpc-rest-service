package echo

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/bwolf1/grpc-rest-kubernetes/proto"
)

func TestServer_Echo(t *testing.T) {
	type fields struct {
		UnimplementedEchoerServer pb.UnimplementedEchoerServer
	}
	type args struct {
		ctx context.Context
		in  *pb.EchoRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.EchoResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				UnimplementedEchoerServer: tt.fields.UnimplementedEchoerServer,
			}
			got, err := s.Echo(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.Echo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.Echo() = %v, want %v", got, tt.want)
			}
		})
	}
}
