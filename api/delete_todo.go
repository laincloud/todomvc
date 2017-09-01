package api

import (
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// DeleteTodo 删除 Todo
func DeleteTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params, db *xorm.Engine, logger *zap.Logger) {
}
