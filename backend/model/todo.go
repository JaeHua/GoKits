package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title     string `gorm:"type:varchar(20);not null"`
	Status    bool   `gorm:"not null;default:false"`
	Telephone string // 添加 Telephone 字段
}
