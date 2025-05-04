
package main

import (
    "fmt"
    "os"
    "text/template"
)

const modelTemplate = 
`
package models

type {{.Name}} struct {
    ID uint gorm:"primaryKey"
    Nom string
}
`

const handlerTemplate = 
`
package handlers

import (
    "github.com/gin-gonic/gin"
)

func Get{{.Name}}(c *gin.Context) {
    c.JSON(200, gin.H{"message": "{{.Name}}"})
}
`

const migrationTemplate = 
`
package migrations

import "gorm.io/gorm"

type {{.Name}} struct {
    ID uint gorm:"primaryKey"
    Nom string
}

func Migrate{{.Name}}(db *gorm.DB) {
    db.AutoMigrate(&{{.Name}}{})
}
`

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
        file, _ := os.Create("../../internal/models/" + name + ".go")
        defer file.Close()
        tmpl.Execute(file, struct{ Name string }{name})
        fmt.Printf("✅ Modèle '%s' créé dans internal/models/\n", name)

    case "make:handler":
        tmpl, _ := template.New("handler").Parse(handlerTemplate)
        file, _ := os.Create("../../internal/handlers/" + name + "_handler.go")
        defer file.Close()
        tmpl.Execute(file, struct{ Name string }{name})
        fmt.Printf("✅ Handler '%s' créé dans internal/handlers/\n", name)

    case "make:migration":
        tmpl, _ := template.New("migration").Parse(migrationTemplate)
        file, _ := os.Create("../../migrations/migrate_" + name + ".go")
        defer file.Close()
        tmpl.Execute(file, struct{ Name string }{name})
        fmt.Printf("✅ Migration '%s' créée dans migrations/\n", name)

    default:
        fmt.Println("Commande inconnue")
    }
}
