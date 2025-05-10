package config

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv" // 🔥 Ajouté ici
    "gopkg.in/yaml.v3"
)

type Config struct {
    Server struct {
        Port int `yaml:"port"`
    } `yaml:"server"`
    JWTSecret   string `yaml:"jwt_secret"`
    LogLevel    string `yaml:"log_level"`
    DBHost      string
    DBUser      string
    DBPassword  string
    DBName      string
}

func LoadConfig() Config {
    // 🔥 Charge le fichier .env à la racine du projet
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("⚠️ Impossible de charger le fichier .env")
    }

    var cfg Config

    // Charger config.yaml
    data, err := os.ReadFile("config/config.yaml")
    if err != nil {
        log.Fatalf("Erreur lecture config.yaml: %v", err)
    }
    yaml.Unmarshal(data, &cfg)

    // Charger .env dans la config
    cfg.DBHost = os.Getenv("DB_HOST")
    cfg.DBUser = os.Getenv("DB_USER")
    cfg.DBPassword = os.Getenv("DB_PASSWORD")
    cfg.DBName = os.Getenv("DB_NAME")
    cfg.JWTSecret = os.Getenv("JWT_SECRET")

    // Vérifier que JWT_SECRET est bien présent
    if cfg.JWTSecret == "" {
        log.Fatal("JWT_SECRET non défini dans .env")
    }

    // Charger le port depuis .env ou config.yaml
    portStr := os.Getenv("APP_PORT")
    fmt.Sscanf(portStr, "%d", &cfg.Server.Port)

    return cfg
}