package services

import (
	HolidayRequest "api/internal/modules/holiday/requests"
	HolidayResponse "api/internal/modules/holiday/responses"
)

type HolidayServiceInterface interface {
	GetHolidays(limit int) (HolidayResponse.HolidaysResponse, error)
	FindHoliday(id uint) (HolidayResponse.HolidayResponse, error)
	CreateHoliday(payload HolidayRequest.HolidayCreatePayload) (HolidayResponse.HolidayResponse, error)
	UpdateHoliday(id uint, payload HolidayRequest.HolidayUpdatePayload) (HolidayResponse.HolidayResponse, error)
	DeleteHoliday(id uint) error
}
