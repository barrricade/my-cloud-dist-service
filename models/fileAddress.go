// Package models 定义数据库 model
package models

import (
	"gorm.io/gorm"
)

type FileAddress struct {
	gorm.Model
	Fid     File   `gorm:"foreignKey:ID" json:"fid,omitempty"`
	Pid     uint   `json:"pid" example:"1"`
	Cid     uint   `json:"cid" example:"2"`
	DirName string `json:"dir_name" example:"新建文件夹"`
}
