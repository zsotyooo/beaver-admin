package models

import (
	UserModels "api/internal/user/models"
	"errors"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Holiday struct {
	gorm.Model
	FromDate     datatypes.Date
	ToDate       datatypes.Date
	Comment      string
	AdminComment string
	Urgent       bool
	Accepted     bool
	UserID       uint `gorm:"foreignKey"`
	User         UserModels.User
}

func (holiday *Holiday) BeforeSave(tx *gorm.DB) (err error) {
	if time.Time(holiday.FromDate).After(time.Time(holiday.ToDate)) {
		err = errors.New("FromDate cannot be later than ToDate")
	}
	return
}
