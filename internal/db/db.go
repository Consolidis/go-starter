package db

import (
    "fmt"
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"

    "elysionne/internal/config"
)

func ConnectDB(cfg config.Config) *gorm.DB {
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        cfg.DBUser,
        cfg.DBPassword,
        cfg.DBHost,
        cfg.DBName,
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Impossible de se connecter à la base : %v", err)
    }

    log.Println("✅ Connecté à la base MySQL")
    return db
}