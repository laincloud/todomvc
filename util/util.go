package util

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"go.uber.org/zap"
)

type key int

const requestIDKey key = 0

// NewContext 将 requestID 写入 ctx
func NewContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}

// FromContext 返回携带 RequestID 信息的 logger
func FromContext(ctx context.Context, logger *zap.Logger) *zap.Logger {
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return logger
	}

	return logger.With(zap.String("RequestID", requestID))
}

// NewRequestID 获取新的 RequestID
func NewRequestID() string {
	bs := make([]byte, 16)
	if _, err := rand.Read(bs); err != nil {
		return "00000000000000000000000000000000"
	}

	return hex.EncodeToString(bs)
}
