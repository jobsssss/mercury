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
		migrator.DropColumn(&user.User{}, "City")
		migrator.DropColumn(&user.User{}, "Introduction")
		migrator.DropColumn(&user.User{}, "Avatar")
	}

	migrate.Add("2023_10_09_100842_add_fields_to_user", up, down)
}
