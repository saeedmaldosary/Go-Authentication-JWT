package initializers

import "github.com/saeedmaldosary/Go-Authentication-JWT/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
