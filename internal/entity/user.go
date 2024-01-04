package entity

import (
	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model         // ID,CreateAt,UpdateAt,DeletedAt
	Email       string `gorm:"type:varchar(100);unique_index"` //唯一索引
	PhoneNumber string `gorm:"type:varchar(20);unique"`
	UserName    string `gorm:"type:varchar(10)"`
	Gender      uint32 `gorm:"type:tinyint(3);not null"`
	Password    string `gorm:"type:varchar(20)"`
	UserType    uint32 `gorm:"type:tinyint(3);index"`
	Extra       string `gorm:"type:text"`
}

