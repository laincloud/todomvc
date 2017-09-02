package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-xorm/xorm"
	"github.com/laincloud/todomvc/gen/models"
	"github.com/laincloud/todomvc/gen/restapi/operations/ping"
	"go.uber.org/zap"
)

// Ping 用于健康检查
func Ping(_ ping.PingParams, _ *xorm.Engine, logger *zap.Logger) middleware.Responder {
	hostname, err := os.Hostname()
	if err != nil {
		logger.Error("os.Hostname() failed.", zap.Error(err))
		message := err.Error()
		return ping.NewPingDefault(http.StatusInternalServerError).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}

	return ping.NewPingOK().WithPayload(fmt.Sprintf("%s is OK.", hostname))
}
