package main

import (
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"

    "elysionne/internal/config"
    "elysionne/internal/db"
    "elysionne/internal/handlers"
    "elysionne/internal/middleware"
)

func main() {
    cfg := config.LoadConfig()
    database := db.ConnectDB(cfg)

    r := gin.Default()

    // Middleware global pour passer DB et config aux handlers
    r.Use(func(c *gin.Context) {
        c.Set("db", database)
        c.Set("config", cfg)
        c.Next()
    })

    // Routes publiques
    r.POST("/register", handlers.Register)
    r.POST("/login", handlers.Login)

    // Routes protégées
    protected := r.Group("/api")
    protected.Use(middleware.AuthMiddleware(cfg))
    {
        protected.GET("/secure", func(c *gin.Context) {
            userID, _ := c.Get("userID")
            c.JSON(http.StatusOK, gin.H{"message": "Accès autorisé", "user_id": userID})
        })
    }

    port := fmt.Sprintf(":%d", cfg.Server.Port)
    fmt.Printf("🚀 Serveur démarré sur http://localhost%s\n", port)
    r.Run(port)
}