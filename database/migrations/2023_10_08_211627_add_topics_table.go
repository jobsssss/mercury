package migrations

import (
	"database/sql"
	"mercury/app/models"
	"mercury/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		models.BaseModel
	}
	type Category struct {
		models.BaseModel
	}

	type Topic struct {
		models.BaseModel
		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:longtext;notnull"`
		UserId     int64  `gorm:"type:bigint;not null;index"`
		CategoryID int64  `gorm:"type:bigint;not null;index"`

		// 会创建 user_id 和 category_id 外键的约束
		User     User
		Category Category
		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Topic{})
	}

	migrate.Add("2023_10_08_211627_add_topics_table", up, down)
}
