package api

// SuccessResponse represents a generic success response
// @name SuccessResponse
type SuccessResponse struct {
	Message string `json:"message" example:"Success"`
}

// ErrorResponse represents a generic error response
// @name ErrorResponse
type ErrorResponse struct {
	Error string `json:"error" example:"Error"`
}

func NewSuccessResponse(message string) SuccessResponse {
	return SuccessResponse{message}
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{Error: err.Error()}
}
