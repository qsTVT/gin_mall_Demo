package model

import "gorm.io/gorm"

// 用户
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string `gorm:"unique"`
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string
	Money          string
}
