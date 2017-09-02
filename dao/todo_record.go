package dao

import (
	"time"

	"github.com/laincloud/todomvc/gen/models"
)

// TodoRecord 表示数据库中存储的 todo 记录
type TodoRecord struct {
	ID        int64     `xorm:"id pk BIGINT autoincr notnull" json:"id"`
	Title     string    `xorm:"title VARCHAR(255) notnull" json:"title"`
	Done      bool      `xorm:"done BOOL notnull" json:"done"`
	CreatedAt time.Time `xorm:"created_at DATETIME notnull" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at DATETIME notnull" json:"updated_at"`
}

// NewTodoRecord 返回初始化后的 Todo
func NewTodoRecord(todo models.Todo) *TodoRecord {
	now := time.Now()
	return &TodoRecord{
		Title:     *todo.Title,
		Done:      *todo.Done,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// TableName 返回 Todo 在数据库里的表名
func (t TodoRecord) TableName() string {
	return "todo_record"
}
