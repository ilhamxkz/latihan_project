package models

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "root:ilham@tcp(localhost:3306)/smkn1golang?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database: " + err.Error())
    }

    DB = db

    // AutoMigrate jika perlu
    DB.AutoMigrate(&User{})
    DB.AutoMigrate(&Role{})
    DB.AutoMigrate(&Menu{})
    DB.AutoMigrate(&Akses{})
    fmt.Println("Database connected")
}

