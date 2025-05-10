package main

import (
    "elysionne/internal/db"
    "elysionne/internal/config"
    "elysionne/internal/models"
    "log"
)

func main() {
    cfg := config.LoadConfig()
    database := db.ConnectDB(cfg)

    log.Println("🚀 Démarrage des migrations...")
    database.AutoMigrate(&models.User{})
    log.Println("✅ Table User migrée")
}