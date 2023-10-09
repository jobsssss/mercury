package migrations

import (
	"database/sql"
	"mercury/app/models/link"
	"mercury/pkg/migrate"

	"gorm.io/gorm"
)

func init() {
	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&link.Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&link.Link{})
	}

	migrate.Add("2023_10_09_085453_add_links_table", up, down)
}
