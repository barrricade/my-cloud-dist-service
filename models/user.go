// Package models 定义数据库 model
package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	NickName  string    `json:"name" example:"晓晓"`
	Account   string    `json:"account" example:"q656141982"`
	Password  string    `json:"password" example:"w11111"`
	Phone     string    `json:"phone" example:"13730884444"`
	Email     string    `json:"email" example:"77677@qq.com"`
	Avatar    string    `json:"avatar"`
	LastLogin time.Time `json:"last_login" example:"2021-10-11 20:11:22"`
}
