package api

import (
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// GetTodo 获取 Todo
func GetTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params, db *xorm.Engine, logger *zap.Logger) {
}
