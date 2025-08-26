package utils

import "github.com/labstack/echo/v4"

type HttpResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

func NewErrorResponse(err error) *HttpResponse {
	e := &HttpResponse{}
	e.Status = "error"
	e.Data = nil
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Message = v.Message
	default:
		e.Message = v.Error()
	}
	return e
}

func NewSuccessResponse(data interface{}) *HttpResponse {
	return &HttpResponse{
		Status:  "success",
		Data:    data,
		Message: nil,
	}
}
