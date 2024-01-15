package repositories

import (
	HolidayModel "api/internal/modules/holiday/models"
	"api/packages/database"

	"gorm.io/gorm"
)

type HolidayRepository struct {
	DB *gorm.DB
}

func New() *HolidayRepository {
	return &HolidayRepository{
		DB: database.Connection(),
	}
}

func (holidayRepository *HolidayRepository) List(limit int) (holidays []HolidayModel.Holiday, err error) {
	err = holidayRepository.DB.Limit(limit).Order("created_at DESC").Find(&holidays).Error
	return
}

func (HolidayRepository *HolidayRepository) Find(id uint) (holiday HolidayModel.Holiday, err error) {
	err = HolidayRepository.DB.First(&holiday, id).Error
	return
}

func (holidayRepository *HolidayRepository) Create(holiday HolidayModel.Holiday) (newHoliday HolidayModel.Holiday, err error) {
	err = holidayRepository.DB.Create(&holiday).Scan(&newHoliday).Error
	return
}

func (holidayRepository *HolidayRepository) Update(holiday HolidayModel.Holiday, fields map[string]interface{}) (HolidayModel.Holiday, error) {
	err := holidayRepository.DB.Model(&holiday).Updates(fields).Error
	return holiday, err
}

func (holidayRepository *HolidayRepository) Delete(id uint) error {
	err := holidayRepository.DB.Delete(&HolidayModel.Holiday{}, id).Error
	return err
}
