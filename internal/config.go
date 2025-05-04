package config

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
}

func LoadConfig() Config {
    var cfg Config

    // Charger config.yaml
    data, err := os.ReadFile("config/config.yaml")
    if err != nil {
        log.Fatalf("Erreur lecture config.yaml: %v", err)
    }
    yaml.Unmarshal(data, &cfg)

    // Si le port est défini dans .env, on le priorise
    portStr := os.Getenv("APP_PORT")
    fmt.Sscanf(portStr, "%d", &cfg.Server.Port)

    return cfg
}