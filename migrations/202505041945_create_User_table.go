
package migrations

import (
    "gorm.io/gorm"
    "elysionne/internal/models"
)

func init() {
    Migrations = append(Migrations, &Migration{
        ID: "202505041945_create_User_table",
        Migrate: func(db *gorm.DB) error {
            return db.AutoMigrate(&models.User{})
        },
        Rollback: func(db *gorm.DB) error {
            return db.Migrator().DropTable(&models.User{})
        },
    })
}
