// Package models 定义数据库 model
package models

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Uid      User   `gorm:"foreignKey:ID" json:"uid,omitempty"`
	FileName string `json:"file_name" example:"sdwdd.jpg"`
	Status   string `json:"status" example:"disable"`
	Public   bool   `json:"public" example:"true"`
	Cover    string `json:"cover"`
}
