package domain

type Response struct {
	Message string `json:"message"`
}

type CustomError struct {
	Message    string
	StatusCode int
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func NewCustomError(message string, statusCode int) *CustomError {
	return &CustomError{
		Message:    message,
		StatusCode: statusCode,
	}
}
