package database

import (
    // _ "github.com/go-sql-driver/mysql"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "fmt"
)

var Eloquent *gorm.DB

func init() {
    var err error
    dsn := "root:root@tcp(127.0.0.1:3306)/wallet_service?charset=utf8mb4&parseTime=True&loc=Local"
    fmt.Printf(dsn)
    Eloquent, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Printf("mysql connect error %v", err)
    }

    if Eloquent.Error != nil {
        fmt.Printf("database error %v", Eloquent.Error)
    }
}