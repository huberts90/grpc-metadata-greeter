package grpc

import (
	"context"
	"net"

	"github.com/huberts90/grpc-metadata-greeter/api"
	"google.golang.org/grpc"
)

type GreeterService struct {
	api.UnimplementedGreeterServer
}

type Server struct {
	grpcServer *grpc.Server
}

func NewServer() *Server {
	return &Server{
		grpcServer: grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				ContextInterceptor(),
			)),
	}
}

func (srv *Server) Serve(port string) error {
	// Open a port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	api.RegisterGreeterServer(srv.grpcServer, &GreeterService{})
	return srv.grpcServer.Serve(listener)
}

func (s *GreeterService) SayHello(ctx context.Context, in *api.HelloRequest) (*api.HelloReply, error) {
	return &api.HelloReply{
		UserAgent:          ctx.Value(USER_AGENT).(string),
		UserAgentLowerCase: ctx.Value(USER_AGENT_LOWER_CASE).(string),
		UserAgentWithX:     ctx.Value(USER_AGENT_WITH_X).(string),
	}, nil
}
