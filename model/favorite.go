package model

import "gorm.io/gorm"

// 收藏
type Favorite struct {
	gorm.Model
	User      User    `gorm:"Foreignkey:UserId"`
	UserId    uint    `gorm:"not null"`
	Product   Product `gorm:"Foreignkey:ProductId"`
	ProductId uint    `gorm:"not null"`
	Boss      User    `gorm:"Foreignkey:BossId"`
	BossId    uint    `gorm:"not null"`
}
