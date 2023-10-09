package migrations

import (
    "database/sql"
    "mercury/app/models/user"
    "mercury/pkg/migrate"

    "gorm.io/gorm"
)

func init() {

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&user.User{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&user.User{})
    }

    migrate.Add("2023_10_08_125442_add_users_table", up, down)
}