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
	Title      string            `json:"title,omitempty"`
	Body       string            `json:"body,omitempty"`
	UserId     int64             `json:"user_id,omitempty"`
	CategoryID int64             `json:"category_id,omitempty"`
	User       user.User         `json:"user"`
	Category   category.Category `json:"category"`
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
