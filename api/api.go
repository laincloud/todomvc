package api

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type handlerFunc func(http.ResponseWriter, *http.Request, httprouter.Params, *xorm.Engine, *zap.Logger)

// Handle 处理 HTTP 请求
func Handle(f handlerFunc, db *xorm.Engine, logger *zap.Logger) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		newLogger := logger.With(zap.String("RequestID", newRequestID()))
		newLogger.Info("Receive a request.",
			zap.String("URL", r.URL.String()),
			zap.String("Method", r.Method),
			zap.Any("Header", r.Header),
			zap.String("RemoteAddr", r.RemoteAddr),
		)

		f(w, r, ps, db, newLogger)

		newLogger.Info("Response has been sent.",
			zap.String("URL", r.URL.String()),
			zap.String("Method", r.Method),
			zap.Any("Header", r.Header),
			zap.String("RemoteAddr", r.RemoteAddr),
		)
	}
}

func newRequestID() string {
	bs := make([]byte, 16)
	if _, err := rand.Read(bs); err != nil {
		return "00000000000000000000000000000000"
	}

	return hex.EncodeToString(bs)
}
