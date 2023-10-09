package migrations

import (
	"database/sql"
	"mercury/app/models/category"
	"mercury/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&category.Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&category.Category{})
	}

	migrate.Add("2023_10_08_190216_add_categories_table", up, down)
}
