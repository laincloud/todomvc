package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-xorm/xorm"
	"github.com/laincloud/todomvc/dao"
	"github.com/laincloud/todomvc/gen/models"
	"github.com/laincloud/todomvc/gen/restapi/operations/todo"
	"go.uber.org/zap"
)

// GetTodo 获取一个 todo 项
func GetTodo(params todo.GetTodoParams, db *xorm.Engine, logger *zap.Logger) middleware.Responder {
	todoID, err := strconv.ParseInt(params.ID, 10, 64)
	if err != nil {
		message := err.Error()
		return todo.NewGetTodoDefault(http.StatusBadRequest).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}

	todoRecord := new(dao.TodoRecord)
	ok, err := db.ID(todoID).Get(todoRecord)
	if err != nil {
		message := err.Error()
		return todo.NewGetTodoDefault(http.StatusInternalServerError).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}

	if !ok {
		message := fmt.Sprintf("%d does not exist in database", todoID)
		return todo.NewGetTodoDefault(http.StatusBadRequest).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}

	logger.Info("db.ID().Get() succeed.",
		zap.Any("TodoRecord", todoRecord),
	)
	return todo.NewGetTodoOK().WithPayload(
		&models.Todo{
			ID:    todoRecord.ID,
			Title: &todoRecord.Title,
			Done:  &todoRecord.Done,
		},
	)
}
