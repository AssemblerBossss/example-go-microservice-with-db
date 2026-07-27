package httpapi

import "context"

const requestIDHeader = "X-Request-ID"

// WithRequestIDContext adds request ID to context and returns new context.
func WithRequestIDContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// RequestIDFromContext extracts request ID from context. Returns empty string if not found.
func RequestIDFromContext(ctx context.Context) string {
	value, _ := ctx.Value(requestIDKey).(string)
	return value
}
