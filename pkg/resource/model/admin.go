package model

import (
	"time"

	"gorm.io/gorm"
)

// 字段
type Admin struct {
	Id            int    `gorm:"autoIncrement"`
	Username      string `gorm:"size:20;index:admins_username_unique,unique;not null"`
	Nickname      string `gorm:"size:200;not null"`
	Sex           int    `gorm:"size:4;not null;default:1"`
	Email         string `gorm:"size:50;index:admins_email_unique,unique;not null"`
	Phone         string `gorm:"size:11;index:admins_phone_unique,unique;not null"`
	Password      string `gorm:"size:255;not null"`
	Avatar        string
	LastLoginIp   string `gorm:"size:255"`
	LastLoginTime time.Time
	Status        int `gorm:"size:1;not null;default:1"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
