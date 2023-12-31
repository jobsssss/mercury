//Package link 模型
package link

import (
	"mercury/app/models"
	"mercury/pkg/database"
)

type Link struct {
	models.BaseModel
	Name string `gorm:"type:varchar(255);not null"`
	URL  string `gorm:"type:varchar(255);not null"`
	models.CommonTimestampsField
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

func (link *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}

func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}
