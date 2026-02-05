package db

import "gorm.io/gorm"

type UserModel struct {
    gorm.Model
    Username string `gorm:"unique;not null"`
    Password string `gorm:"not null"`
}

func (UserModel) TableName() string {
    return "users"
}
