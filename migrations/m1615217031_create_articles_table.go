package migrations

import (
	"course-go/models"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func m1615217031CreateArticlesTable() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "1615217031",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&models.Article{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTable("articles").Error
		},
	}
}
