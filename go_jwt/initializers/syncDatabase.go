package initializers

import (
	"github.com/master-of-none/go_jwt/models"
)

func SyncDB() {
	DB.AutoMigrate(&models.User{})
}
