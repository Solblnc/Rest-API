package database

import (
	"github.com/Solblnc/Rest-API/internal/coment"
	"github.com/jinzhu/gorm"
)

func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&coment.Comment{}); result.Error != nil {
		return result.Error
	}

	return nil
}
