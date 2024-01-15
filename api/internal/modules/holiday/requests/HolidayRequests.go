package holiday

type HolidayCreatePayload struct {
	FromDate     string `json:"from_date" binding:"required"`
	ToDate       string `json:"to_date" binding:"required"`
	Comment      string `json:"comment" binding:"required"`
	AdminComment string `json:"admin_comment"`
	Urgent       *bool  `json:"urgent,omitempty"`
	Accepted     *bool  `json:"accepted,omitempty"`
	UserID       uint   `json:"user_id" binding:"required"`
}

type HolidayUpdatePayload struct {
	FromDate     *string `json:"from_date,omitempty"`
	ToDate       *string `json:"to_date,omitempty"`
	Comment      *string `json:"comment,omitempty"`
	AdminComment *string `json:"admin_comment,omitempty"`
	Urgent       *bool   `json:"urgent,omitempty,omitempty"`
	Accepted     *bool   `json:"accepted,omitempty,omitempty"`
}
