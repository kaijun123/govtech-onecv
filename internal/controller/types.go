package controller

type Request struct {
	Student  string `json:"student,omitempty"`
	Students string `json:"students,omitempty"`
	Teacher  string `json:"teacher,omitempty"`
}

type Response struct {
	Students   []string `json:"students,omitempty"`
	Recipients []string `json:"recipients,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Message: message,
	}
}

func NewSuccessResponse(status string) SuccessResponse {
	if status != "" {
		return SuccessResponse{
			Status: status,
		}
	} else {
		return SuccessResponse{
			Status: "Success",
		}
	}
}
