package handler

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-xorm/xorm"
	"github.com/laincloud/todomvc/dao"
	"github.com/laincloud/todomvc/gen/models"
	"github.com/laincloud/todomvc/gen/restapi/operations/todo"
	"go.uber.org/zap"
)

// CreateTodo 创建一个 todo 项
func CreateTodo(params todo.CreateTodoParams, db *xorm.Engine, logger *zap.Logger) middleware.Responder {
	todoRecord := dao.NewTodoRecord(*params.Body)
	affected, err := db.InsertOne(todoRecord)
	if err != nil {
		message := err.Error()
		return todo.NewCreateTodoDefault(http.StatusInternalServerError).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}

	logger.Info("db.InsertOne() succeed.",
		zap.Any("TodoRecord", todoRecord),
		zap.Int64("Affected", affected),
	)
	return todo.NewCreateTodoCreated().WithPayload(
		&models.Todo{
			ID:    todoRecord.ID,
			Title: &todoRecord.Title,
			Done:  &todoRecord.Done,
		},
	)
}
