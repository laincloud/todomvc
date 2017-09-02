package handler

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-xorm/xorm"
	"github.com/laincloud/todomvc/dao"
	"github.com/laincloud/todomvc/gen/models"
	"github.com/laincloud/todomvc/gen/restapi/operations/todo"
	"go.uber.org/zap"
)

// DeleteTodo 删除一个 todo 项
func DeleteTodo(params todo.DeleteTodoParams, db *xorm.Engine, logger *zap.Logger) middleware.Responder {
	todoID, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		message := err.Error()
		return todo.NewDeleteTodoDefault(http.StatusBadRequest).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}

	todoRecord := new(dao.TodoRecord)
	affected, err := db.ID(todoID).Delete(todoRecord)
	if err != nil {
		message := err.Error()
		return todo.NewDeleteTodoDefault(http.StatusInternalServerError).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}
	logger.Info("db.ID().Delete() succeed.",
		zap.Int64("TodoID", todoID),
		zap.Int64("Affected", affected),
	)
	return todo.NewDeleteTodoNoContent()
}
