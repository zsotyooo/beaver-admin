package services

import (
	HolidayRequest "api/internal/holiday/requests"
	HolidayResponse "api/internal/holiday/responses"
)

type HolidayServiceInterface interface {
	GetHolidays(limit int) (HolidayResponse.HolidaysResponse, error)
	FindHoliday(id uint) (HolidayResponse.HolidayResponse, error)
	CreateHoliday(payload HolidayRequest.HolidayCreatePayload) (HolidayResponse.HolidayResponse, error)
	UpdateHoliday(id uint, payload HolidayRequest.HolidayUpdatePayload) (HolidayResponse.HolidayResponse, error)
	DeleteHoliday(id uint) error
}
