package middleware

import (
    // "fmt"
    "net/http"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "elysionne/internal/config"
)

func AuthMiddleware(cfg config.Config) gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token manquant"})
            return
        }

        claims := &jwt.StandardClaims{}
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(cfg.JWTSecret), nil
        })

        if err != nil || !token.Valid {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token invalide", "details": err.Error()})
            return
        }

        c.Set("userID", claims.Issuer)
        c.Next()
    }
}