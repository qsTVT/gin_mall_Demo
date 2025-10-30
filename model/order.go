package model

import "gorm.io/gorm"

// 订单
type Order struct {
	gorm.Model
	UserId    uint   `gorm:"not null"`
	ProductId uint   `gorm:"not null"`
	BossId    uint   `gorm:"not null"`
	AddressId string `gorm:"not null"`
	Num       int
	OrderNum  uint64
	Type      uint //1未支付 2已支付
	Money     float64
}
