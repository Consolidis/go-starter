package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "elysionne/internal"
)

func main() {
    cfg := config.LoadConfig()

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Bienvenue dans ton nouveau projet Go !"})
    })

    port := fmt.Sprintf(":%d", cfg.Server.Port)
    fmt.Printf("🚀 Serveur démarré sur http://localhost%s\n", port)
    r.Run(port)
}