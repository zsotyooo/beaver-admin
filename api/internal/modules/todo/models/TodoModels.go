package models

import (
	// UserModels "api/internal/modules/user/models"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title string
	Done  bool
	// UserID uint
	// User   UserModels.User
}
