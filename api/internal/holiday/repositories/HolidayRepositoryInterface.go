package repositories

import (
	HolidayModel "api/internal/holiday/models"
)

type HolidayRepositoryInterface interface {
	List(limit int) ([]HolidayModel.Holiday, error)
	Find(id uint) (HolidayModel.Holiday, error)
	Create(holiday HolidayModel.Holiday) (HolidayModel.Holiday, error)
	Update(holiday HolidayModel.Holiday, fields map[string]interface{}) (HolidayModel.Holiday, error)
	Delete(id uint) error
}
