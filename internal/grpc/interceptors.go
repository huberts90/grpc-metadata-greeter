package grpc

import (
	"context"
	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

const (
	USER_AGENT            = "User-Agent"
	USER_AGENT_LOWER_CASE = "user-agent"
	USER_AGENT_WITH_X     = "x-user-agent"
)

func ContextInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		userAgent := getValueFromIncomingMetadata(ctx, USER_AGENT)
		userAgentLowerCase := getValueFromIncomingMetadata(ctx, USER_AGENT_LOWER_CASE)
		userAgentWithX := getValueFromIncomingMetadata(ctx, USER_AGENT_WITH_X)

		ctx = context.WithValue(ctx, USER_AGENT, userAgent)
		ctx = context.WithValue(ctx, USER_AGENT_LOWER_CASE, userAgentLowerCase)
		ctx = context.WithValue(ctx, USER_AGENT_WITH_X, userAgentWithX)

		resp, err := handler(ctx, req)

		return resp, err
	}
}

func getValueFromIncomingMetadata(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	mdData, ok := md[key]
	if !ok || len(mdData) < 1 {
		return ""
	}

	return mdData[0]
}
