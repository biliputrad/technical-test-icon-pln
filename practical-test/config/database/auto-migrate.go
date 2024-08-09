package database

import "gorm.io/gorm"

func MigrateTables(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&domain.User{}, &domain.Article{})

	return err
}
