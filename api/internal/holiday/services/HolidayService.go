package services

import (
	HolidayModel "api/internal/holiday/models"
	HolidayRepository "api/internal/holiday/repositories"
	HolidayRequest "api/internal/holiday/requests"
	HolidayResponse "api/internal/holiday/responses"
	"errors"
	"time"

	"api/pkg/converters"

	"gorm.io/datatypes"
)

type HolidayService struct {
	holidayRepository HolidayRepository.HolidayRepositoryInterface
}

func New() *HolidayService {
	return &HolidayService{
		holidayRepository: HolidayRepository.New(),
	}
}

func (holidayService *HolidayService) GetHolidays(limit int) (HolidayResponse.HolidaysResponse, error) {
	holidays, err := holidayService.holidayRepository.List(limit)

	if err != nil {
		return HolidayResponse.HolidaysResponse{}, err
	}

	return HolidayResponse.ConvertModelsToResponse(holidays), nil
}

func (holidayService *HolidayService) FindHoliday(id uint) (HolidayResponse.HolidayResponse, error) {
	var response HolidayResponse.HolidayResponse

	holiday, err := holidayService.holidayRepository.Find(id)

	if err != nil {
		return response, err
	}

	if holiday.ID == 0 {
		return response, errors.New("Holiday not found!")
	}

	return HolidayResponse.ConvertModelToResponse(holiday), nil
}

func (holidayService *HolidayService) CreateHoliday(payload HolidayRequest.HolidayCreatePayload) (response HolidayResponse.HolidayResponse, err error) {
	var holiday HolidayModel.Holiday

	fromDate, err := time.Parse(time.DateTime, payload.FromDate)
	if err != nil {
		return
	}

	toDate, err := time.Parse(time.DateTime, payload.ToDate)
	if err != nil {
		return
	}

	holiday.FromDate = datatypes.Date(fromDate)
	holiday.ToDate = datatypes.Date(toDate)
	holiday.Comment = payload.Comment
	holiday.AdminComment = payload.AdminComment
	holiday.Urgent = *payload.Urgent
	holiday.Accepted = *payload.Accepted
	holiday.UserID = payload.UserID

	newHoliday, err := holidayService.holidayRepository.Create(holiday)

	if err != nil {
		return
	}

	if newHoliday.ID == 0 {
		err = errors.New("Error in creating the holiday!")
		return
	}

	response = HolidayResponse.ConvertModelToResponse(newHoliday)
	return
}

func (holidayService *HolidayService) UpdateHoliday(id uint, payload HolidayRequest.HolidayUpdatePayload) (HolidayResponse.HolidayResponse, error) {
	var response HolidayResponse.HolidayResponse
	holiday, err := holidayService.holidayRepository.Find(id)

	if err != nil {
		return response, err
	}

	if holiday.ID == 0 {
		return response, errors.New("Holiday not found!")
	}

	fields, err := converters.StructToMap(payload)

	if err != nil {
		return response, err
	}

	updatedHoliday, err := holidayService.holidayRepository.Update(holiday, fields)

	if err != nil {
		return response, err
	}

	return HolidayResponse.ConvertModelToResponse(updatedHoliday), nil
}

func (holidayService *HolidayService) DeleteHoliday(id uint) error {
	holiday, err := holidayService.holidayRepository.Find(id)

	if err != nil {
		return err
	}

	if holiday.ID == 0 {
		return errors.New("Holiday not found!")
	}

	holidayService.holidayRepository.Delete(id)

	return nil
}
