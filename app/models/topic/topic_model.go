//Package topic 模型
package topic

import (
	"mercury/app/models"
	"mercury/app/models/category"
	"mercury/app/models/user"
	"mercury/pkg/database"
)

type Topic struct {
	models.BaseModel
	Title      string `gorm:"type:varchar(255);not null;index"`
	Body       string `gorm:"type:longtext;notnull"`
	UserID     string `gorm:"type:bigint;not null;index"`
	CategoryID string `gorm:"type:bigint;not null;index"`
	User       user.User
	Category   category.Category
	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
