package models

import "gorm.io/gorm"

type Books struct {
	Id        uint    `gorm:"primery key;authoIncrement" json:"id"`
	Title     *string `json:"title"`
	Author    *string `json:"author"`
	Publisher *string `json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err

}
