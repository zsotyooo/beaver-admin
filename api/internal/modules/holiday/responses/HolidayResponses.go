package responses

import (
	HolidayModel "api/internal/modules/holiday/models"
	UserResponses "api/internal/modules/user/responses"
	"time"

	"github.com/thoas/go-funk"
)

type HolidayResponse struct {
	ID           uint                       `json:"id"`
	FromDate     string                     `json:"from_date"`
	ToDate       string                     `json:"to_date"`
	Comment      string                     `json:"comment"`
	AdminComment string                     `json:"admin_comment"`
	Urgent       bool                       `json:"urgent"`
	Accepted     bool                       `json:"accepted"`
	UserID       uint                       `json:"user_id"`
	User         UserResponses.UserResponse `json:"user"`
	CreatedAt    string                     `json:"createdAt"`
}

type HolidaysResponse struct {
	Data []HolidayResponse `json:"data"`
}

func ConvertModelToResponse(holiday HolidayModel.Holiday) HolidayResponse {
	return HolidayResponse{
		ID:           holiday.ID,
		FromDate:     time.Time(holiday.FromDate).Format(time.RFC3339),
		ToDate:       time.Time(holiday.ToDate).Format(time.RFC3339),
		Comment:      holiday.Comment,
		AdminComment: holiday.AdminComment,
		Urgent:       holiday.Urgent,
		Accepted:     holiday.Accepted,
		UserID:       holiday.UserID,
		User:         UserResponses.ConvertModelToResponse(holiday.User),
		CreatedAt:    holiday.CreatedAt.Format(time.RFC3339),
	}
}

func ConvertModelsToResponse(holidays []HolidayModel.Holiday) HolidaysResponse {
	return HolidaysResponse{
		Data: funk.Map(holidays, func(holiday HolidayModel.Holiday) HolidayResponse {
			return ConvertModelToResponse(holiday)
		}).([]HolidayResponse),
	}
}
