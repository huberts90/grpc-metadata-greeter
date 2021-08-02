package grpc

import (
	"context"
	"fmt"
	"github.com/huberts90/grpc-metadata-greeter/api"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	PORT_CONTEXT_INTERCEPTOR = "9001"
)

func Test_ContextInterceptor(t *testing.T) {
	grpcServerIsRunning := new(sync.WaitGroup)
	grpcServerIsRunning.Add(1)

	go startGrpcServer(t, grpcServerIsRunning, PORT_CONTEXT_INTERCEPTOR)

	grpcServerIsRunning.Wait()
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%s", PORT_CONTEXT_INTERCEPTOR), grpc.WithInsecure())
	require.Nil(t, err)
	require.NotNil(t, conn)

	defer conn.Close()

	// values
	userAgent := "greeter-api-1"
	userAgentLowerCase := "greeter-api-2"
	userAgentWithX := "greeter-api-3"

	// Set metadata
	ctx := metadata.AppendToOutgoingContext(
		context.Background(),
		USER_AGENT, userAgent,
		USER_AGENT_LOWER_CASE, userAgentLowerCase,
		USER_AGENT_WITH_X, userAgentWithX)

	c := api.NewGreeterClient(conn)
	r, err := c.SayHello(ctx, &api.HelloRequest{})
	require.Nil(t, err)

	// Case 1: keys has to in lower-case format: https://github.com/grpc/grpc-go/blob/ea9b7a0a7651baaf43c5403cb83349fffb5162df/metadata/metadata.go#L187
	assert.Equal(t, userAgent, r.UserAgent, fmt.Sprintf("Key: %s", USER_AGENT))

	// Case 2: overwrote by engine
	assert.Equal(t, userAgentLowerCase, r.UserAgentLowerCase, fmt.Sprintf("Key: %s", USER_AGENT_LOWER_CASE))

	// Case 3: success
	assert.Equal(t, userAgentWithX, r.UserAgentWithX, fmt.Sprintf("Key: %s", USER_AGENT_WITH_X))
}

func startGrpcServer(t *testing.T, isServerRunning *sync.WaitGroup, port string) {
	server := NewServer()
	isServerRunning.Done()
	err := server.Serve(port)
	require.Nil(t, err)
}
