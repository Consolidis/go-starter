package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"elysionne/internal/models"
	"elysionne/internal/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("votre-cle-secrete")

type Credentials struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func Register(c *gin.Context) {
    var creds Credentials
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Printf("Erreur de hashage du mot de passe: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de créer l'utilisateur"})
        return
    }

    user := models.User{
        Email:    creds.Email,
        Password: string(hashedPassword),
    }

    db := c.MustGet("db").(*gorm.DB)

    result := db.Create(&user)
    if result.Error != nil {
        log.Printf("Erreur lors de la création de l'utilisateur: %v", result.Error)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'enregistrement", "details": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Utilisateur créé", "user": user})
}

func Login(c *gin.Context) {
    var creds Credentials
    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
        return
    }

    var user models.User
    db := c.MustGet("db").(*gorm.DB)
    db.Where("email = ?", creds.Email).First(&user)

    if user.ID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non trouvé"})
        return
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Mot de passe incorrect"})
        return
    }

    cfg := c.MustGet("config").(config.Config) // ✅ On récupère la config

    expirationTime := time.Now().Add(5 * time.Hour)
    claims := &jwt.StandardClaims{
        ExpiresAt: expirationTime.Unix(),
        Issuer:    fmt.Sprint(user.ID),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte(cfg.JWTSecret)) // ✅ Avec clé depuis .env
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la génération du token", "details": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}