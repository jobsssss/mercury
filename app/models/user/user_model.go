package user

import (
	"mercury/app/models"
	"mercury/pkg/database"
	"mercury/pkg/hash"
)

type User struct {
	models.BaseModel

	Name         string `gorm:"type:varchar(255);not null;index"`
	Email        string `gorm:"type:varchar(255);index;default:null"`
	Phone        string `gorm:"type:varchar(20);index;default:null"`
	Password     string `gorm:"type:varchar(255)"`
	City         string `gorm:"type:varchar(10);"`
	Introduction string `gorm:"type:varchar(255);"`
	Avatar       string `gorm:"type:varchar(255);default:null"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
