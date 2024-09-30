package grpc

import (
	"context"

	"github.com/jamm3e3333/quiz-app/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func NewLogInterceptor(lg logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		md.Set("authorization", "***")

		reqKeyToValue := map[string]any{
			"method":       info.FullMethod,
			"user_agent":   valueForMD(md, "user-agent"),
			"request_body": req,
		}
		resp, err := handler(ctx, req)

		resKeyToValue := map[string]any{
			"status": codes.OK,
		}
		if resp != nil {
			resKeyToValue["response_body"] = resp
		}

		if err != nil {
			var errMsg string
			st, ok := status.FromError(err)
			if !ok {
				errMsg = err.Error()
				resKeyToValue["status"] = codes.Internal
			} else {
				resKeyToValue["status"] = st.Code()
				errMsg = st.Message()
			}
			lg.Error(errMsg)
		}

		lg.InfoWithMetadata("gRPC Request / Response", map[string]any{
			"request":  reqKeyToValue,
			"response": resKeyToValue,
		})

		return resp, err
	}
}

func valueForMD(md metadata.MD, key string) string {
	v := md.Get(key)
	if len(v) > 0 {
		return v[0]
	}

	return ""
}
