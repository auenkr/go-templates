// Package interceptors contains interceptors for the application.
package interceptors

import (
	"context"
	"time"

	"connectrpc.com/connect"
	"go.uber.org/zap"
)

func NewLatencyLoggerInterceptor(logger *zap.Logger) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			start := time.Now()
			res, err := next(ctx, req)
			logger.Info(
				"Request Latency",
				zap.String("method", req.Spec().Procedure),
				zap.String("method", req.HTTPMethod()),
				zap.Duration("latency", time.Since(start)),
			)
			return res, err
		}
	}
}
