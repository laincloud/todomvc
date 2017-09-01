package api

import (
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// GetTodos 获取 Todos
func GetTodos(w http.ResponseWriter, r *http.Request, _ httprouter.Params, db *xorm.Engine, logger *zap.Logger) {
}
