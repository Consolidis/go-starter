
package models

type User struct {
    ID uint `gorm:"primaryKey"`
    Nom string
    Email string `gorm:"unique"`
    Password string
    Ville string
    CodePostal string
    Adresse string
    Numero string
}
