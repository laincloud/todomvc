package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// Ping 用于健康检查
func Ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params, _ *xorm.Engine, logger *zap.Logger) {
	hostname, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err = fmt.Fprintf(w, "%s is OK.\n", hostname); err != nil {
		logger.Error("fmt.Fprintf() failed.",
			zap.String("Hostname", hostname),
			zap.Error(err),
		)
	}
}
