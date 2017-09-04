package handler

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-xorm/xorm"
	"github.com/laincloud/todomvc/config"
	"github.com/laincloud/todomvc/gen/restapi/operations/static"
	"github.com/laincloud/todomvc/util"
	"go.uber.org/zap"
)

// ServeIndexHTML 返回主页
func ServeIndexHTML(_ static.ServeIndexHTMLParams, _ *xorm.Engine, logger *zap.Logger) middleware.Responder {
	return middleware.ResponderFunc(func(w http.ResponseWriter, _ runtime.Producer) {
		util.ServeFile(w, fmt.Sprintf("%s/index.html", config.StaticFileDirectory), logger)
	})
}
