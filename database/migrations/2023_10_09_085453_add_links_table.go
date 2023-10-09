package migrations

import (
	"database/sql"
	"mercury/app/models"
	"mercury/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel
		Name string `gorm:"type:varchar(255);not null"`
		URL  string `gorm:"type:varchar(255);not null"`
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Link{})
	}

	migrate.Add("2023_10_09_085453_add_links_table", up, down)
}
