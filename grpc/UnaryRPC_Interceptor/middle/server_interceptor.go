package middle

import (
	"context"
	"google.golang.org/grpc"
	"log/slog"
)

// LogInterceptor 是一个 gRPC 拦截器，用于记录每个 gRPC 调用的日志信息。
func LogInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// 在调用处理程序之前添加自定义逻辑
	slog.Info("UnaryRPC logInterceptor [before]", "method:", info.FullMethod, "request:", req)

	// 调用实际的 gRPC 方法处理程序
	resp, err = handler(ctx, req)

	// 在调用处理程序之后添加自定义逻辑
	slog.Info("UnaryRPC logInterceptor [end]", "method:", info.FullMethod, "request:", req)

	return resp, err
}
