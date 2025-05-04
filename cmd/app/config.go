package main

import (
    "fmt"
    "log"
    "os"

    "gopkg.in/yaml.v3"
)

type Config struct {
    Server struct {
        Port int `yaml:"port"`
    } `yaml:"server"`
    LogLevel string `yaml:"log_level"`
    DBHost   string
    DBUser   string
    DBPass   string
    DBName   string
}

func LoadConfig() Config {
    var cfg Config

    // Charger config.yaml
    data, err := os.ReadFile("../config/config.yaml")
    if err != nil {
        log.Fatalf("Erreur lecture config.yaml: %v", err)
    }
    yaml.Unmarshal(data, &cfg)

    // Charger .env
    cfg.DBHost = os.Getenv("DB_HOST")
    cfg.DBUser = os.Getenv("DB_USER")
    cfg.DBPass = os.Getenv("DB_PASSWORD")
    cfg.DBName = os.Getenv("DB_NAME")

    fmt.Printf("🔧 Config chargée : %+v\n", cfg)
    return cfg
}