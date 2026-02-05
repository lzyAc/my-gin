package db

import (
    // "gorm.io/gorm"
    "time"
)

// TODO 重构user表； 目前只是测试整个框架
type UserModel struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string    `gorm:"unique;not null"`
    Password  string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// db.UserDB.Create(&u).Error 如果不了解GORM 这里可能会有点疑惑 怎么用到这个users表的
// GORM 会检查每个模型（结构体） 是否实现TableName方法
// 如果实现了，gorm就用你返沪ide字符串作为表名
// 如果没有实现，GORM 会把默认结构体名转成蛇形命名，UserModel => user_models
func (UserModel) TableName() string {
    return "users"
}
