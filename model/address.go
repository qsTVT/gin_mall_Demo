package model

import "gorm.io/gorm"

// 收货地址
type Address struct {
	gorm.Model
	UerID   int    `gorm:"not null"`
	Name    string `gorm:"type:varchar(20);not null"`
	Phone   string `gorm:"type:varchar(11);not null"`
	Address string `gorm:"type:varchar(50);not null"`
}
