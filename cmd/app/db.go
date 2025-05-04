package main

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func ConnectDB(cfg Config) *gorm.DB {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBName)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Impossible de se connecter à la base")
    }
    fmt.Println("✅ Connecté à la base MySQL")
    return db
}