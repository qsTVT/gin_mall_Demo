package model

import "gorm.io/gorm"

// 购物车 Boss是商家
type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"` //商品
	BossId    uint `gorm:"not null"` //商家
	Num       uint `gorm:"not null"`
	MaxNum    uint `gorm:"not null"` //商品数量限额
	Check     bool //是否支付
}
