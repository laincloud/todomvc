package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-xorm/xorm"
	"github.com/laincloud/todomvc/dao"
	"github.com/laincloud/todomvc/gen/models"
	"github.com/laincloud/todomvc/gen/restapi/operations/todo"
	"go.uber.org/zap"
)

// UpdateTodo 更新一个 todo 项
func UpdateTodo(params todo.UpdateTodoParams, db *xorm.Engine, logger *zap.Logger) middleware.Responder {
	todoID, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		message := err.Error()
		return todo.NewDeleteTodoDefault(http.StatusBadRequest).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}

	todoRecord := &dao.TodoRecord{
		Title:     *params.Body.Title,
		Completed: *params.Body.Completed,
		UpdatedAt: time.Now(),
	}
	affected, err := db.ID(todoID).UseBool("completed").Update(todoRecord)
	if err != nil {
		message := err.Error()
		return todo.NewUpdateTodoDefault(http.StatusInternalServerError).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}
	logger.Info("db.ID().Update() succeed.",
		zap.Any("TodoRecord", todoRecord),
		zap.Int64("Affected", affected),
	)
	return todo.NewUpdateTodoNoContent()
}
