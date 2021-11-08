// Package models 定义数据库 model
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName  string    `json:"name" example:"晓晓" binding:"required"`
	Account   string    `json:"account" example:"q656141982" binding:"required" gorm:"unique"`
	Password  string    `json:"password" example:"w11111" binding:"required"`
	Phone     string    `json:"phone" example:"13730884444" binding:"required" gorm:"unique"`
	Email     string    `json:"email" example:"77677@qq.com" gorm:"unique"`
	Avatar    string    `json:"avatar"`
	LastLogin time.Time `json:"last_login" example:"2021-10-11 20:11:22"`
	Status    string    `json:"status" gorm:"default:'able'"`
}
