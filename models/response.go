package models

type HttpError struct {
	Message string `json:"message" default:"Some Error Occurred"`
}

type HttpSuccess struct {
	Message string `json:"message" default:"Success"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewHttpError(message string) *HttpError {
	return &HttpError{Message: message}
}

func NewHttpSuccess(message string) *HttpSuccess {
	return &HttpSuccess{Message: message}
}
