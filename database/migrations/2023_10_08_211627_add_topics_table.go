package migrations

import (
	"database/sql"
	"mercury/app/models/topic"
	"mercury/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&topic.Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&topic.Topic{})
	}

	migrate.Add("2023_10_08_211627_add_topics_table", up, down)
}
