package handler

import (
	"github.com/wudtichaikarun/poc-go-template/util"
	"github.com/wudtichaikarun/poc-go-template/util/apperror"
)

type Response[D any] struct {
	Status  int            `json:"status"`
	Code    *string        `json:"code,omitempty"`
	Message *string        `json:"message,omitempty"`
	Data    *D             `json:"data,omitempty"`
	ErrData map[string]any `json:"errorData,omitempty"`
}

func NewMessageResponse(status int, message string) Response[any] {
	return Response[any]{Status: status, Message: util.ToPointer(message)}
}

func NewMessageResponseOK() Response[any] {
	return Response[any]{Status: 200, Message: util.ToPointer("ok")}
}

func NewDataResponse[T any](data T) Response[T] {
	return Response[T]{Status: 200, Data: util.ToPointer(data)}
}

func NewErrorResponse(err *apperror.AppError) Response[any] {
	return Response[any]{
		Status:  err.HTTPStatusCode,
		Code:    &err.Code,
		Message: util.ToPointer(err.Error()),
		ErrData: err.Data,
	}
}
