package internal

import (
	"context"
	"time"
)

type ContextKey string
const DefaultHttpTimeout = 20 * time.Second

const ContextKeyCorrelationID ContextKey = "correlation-id"

func GetContextValue(ctx context.Context, key ContextKey) string {
	reqID := ctx.Value(key)
	if reqID != nil {
		if ret, ok := reqID.(string); ok {
			return ret
		}
	}
	return ""
}

func NewContextWithTimeOut(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, timeout)
}

func SetContextWithValue(ctx context.Context, key ContextKey, value string) context.Context {
	return context.WithValue(ctx, key, value)
}
