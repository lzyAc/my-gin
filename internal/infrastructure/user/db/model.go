package db

import (
    // "gorm.io/gorm"
    "time"
)

type UserModel struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (UserModel) TableName() string {
    return "users"
}
