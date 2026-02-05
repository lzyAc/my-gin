package db

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var (
    UserDB *gorm.DB
    // OrgDB  *gorm.DB 可扩展
)

func InitDB() {
    UserDB = mustOpen(
        "root:gin123@tcp(127.0.0.1:3306)/db_user?charset=utf8mb4&parseTime=True&loc=Local",
    )

    // OrgDB = mustOpen(
    //     "root:gin123@tcp(127.0.0.1:3306)/db_org?charset=utf8mb4&parseTime=True&loc=Local",
    // )
}

func mustOpen(dsn string) *gorm.DB {
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database:", err)
    }
    return db
}
