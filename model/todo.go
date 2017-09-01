package model

import (
	"time"
)

// Todo 表示待办事项
type Todo struct {
	ID        int64     `xorm:"id pk BIGINT autoincr notnull" json:"id"`
	Title     string    `xorm:"title VARCHAR(255) notnull" json:"title"`
	Done      bool      `xorm:"done BOOL notnull" json:"done"`
	CreatedAt time.Time `xorm:"created_at DATETIME notnull" json:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at DATETIME notnull" json:"updated_at"`
}

// TableName 返回 Todo 在数据库里的表名
func (t Todo) TableName() string {
	return "todo"
}
