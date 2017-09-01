package api

import (
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

// UpdateTodo 更新 Todo
func UpdateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params, db *xorm.Engine, logger *zap.Logger) {
}
