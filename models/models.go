// Package models 定义数据库 model
package models

import (
	"time"
)

// gorm model字段信息
type BaseModel struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
type User struct {
	BaseModel
	Name    string `json:"name" example:"red wine No14"`
	Account int    `json:"account_id" example:"1"`
}
