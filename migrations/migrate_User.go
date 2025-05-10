
package migrations

import "gorm.io/gorm"

type User struct {
    ID uint `gorm:"primaryKey"`
    Nom string
}

func MigrateUser(db *gorm.DB) {
    db.AutoMigrate(&User{})
}
