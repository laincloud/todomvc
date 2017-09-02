package handler

import (
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-xorm/xorm"
	"github.com/laincloud/todomvc/dao"
	"github.com/laincloud/todomvc/gen/models"
	"github.com/laincloud/todomvc/gen/restapi/operations/todo"
	"go.uber.org/zap"
)

// ListTodos 列出 todo 列表
func ListTodos(params todo.ListTodosParams, db *xorm.Engine, logger *zap.Logger) middleware.Responder {
	todoRecords := make([]dao.TodoRecord, 0, *params.Limit)
	if err := db.Where("created_at >= ?", time.Unix(0, *params.Since)).Limit(int(*params.Limit), int(*params.Offset)).Find(&todoRecords); err != nil {
		message := err.Error()
		return todo.NewListTodosDefault(http.StatusInternalServerError).WithPayload(
			&models.Error{
				Message: &message,
			},
		)
	}

	todos := make([]*models.Todo, len(todoRecords))
	for i, record := range todoRecords {
		title := record.Title
		done := record.Done
		todos[i] = &models.Todo{
			ID:    record.ID,
			Title: &title,
			Done:  &done,
		}
	}
	return todo.NewListTodosOK().WithPayload(todos)
}
