package main

import (
    "fmt"
    "os"
    "text/template"
    "time"
)

// Templates
const modelTemplate = `
package models

type {{.Name}} struct {
    ID uint gorm:"primaryKey"
    Nom string
}
`

const handlerTemplate = `
package handlers

import (
    "github.com/gin-gonic/gin"
)

func Get{{.Name}}(c *gin.Context) {
    c.JSON(200, gin.H{"message": "{{.Name}}"})
}
`

// Migration versionnée avec ID + Migrate + Rollback
const migrationTemplate = `
package migrations

import (
    "gorm.io/gorm"
    "elysionne/internal/models"
)

func init() {
    Migrations = append(Migrations, &Migration{
        ID: "{{.ID}}",
        Migrate: func(db *gorm.DB) error {
            return db.AutoMigrate(&models.{{.Name}}{})
        },
        Rollback: func(db *gorm.DB) error {
            return db.Migrator().DropTable(&models.{{.Name}}{})
        },
    })
}
`

// MigrationData pour remplir le template
type MigrationData struct {
    ID   string
    Name string
}

// Fonction utilitaire pour obtenir le répertoire courant
func getProjectRoot() string {
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    return dir
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run generator/generator.go make:<type> <name>")
        return
    }

    cmd := os.Args[1]
    name := os.Args[2]

    switch cmd {
    case "make:model":
        tmpl, _ := template.New("model").Parse(modelTemplate)
        filePath := fmt.Sprintf("%s/internal/models/%s.go", getProjectRoot(), name)
        file, _ := os.Create(filePath)
        defer file.Close()
        tmpl.Execute(file, struct{ Name string }{name})
        fmt.Printf("✅ Modèle '%s' créé dans %s\n", name, filePath)

    case "make:handler":
        tmpl, _ := template.New("handler").Parse(handlerTemplate)
        filePath := fmt.Sprintf("%s/internal/handlers/%s_handler.go", getProjectRoot(), name)
        file, _ := os.Create(filePath)
        defer file.Close()
        tmpl.Execute(file, struct{ Name string }{name})
        fmt.Printf("✅ Handler '%s' créé dans %s\n", name, filePath)

    case "make:migration":
        now := time.Now().Format("200601021504")
        id := now + "_create_" + name + "_table"

        data := MigrationData{
            ID:   id,
            Name: name,
        }

        tmpl, _ := template.New("migration").Parse(migrationTemplate)
        filePath := fmt.Sprintf("%s/migrations/%s_create_%s_table.go", getProjectRoot(), now, name)
        file, _ := os.Create(filePath)
        defer file.Close()
        tmpl.Execute(file, data)
        fmt.Printf("✅ Fichier de migration généré : %s\n", filePath)

    default:
        fmt.Println("Commande inconnue")
    }
}