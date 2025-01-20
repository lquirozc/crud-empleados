package models

type Result[T any] struct {
	IsSuccess bool   `json:"isSuccess"`
	Message   string `json:"message"`
	Data      *T     `json:"data"`
}

func Success[T any](data T) Result[T] {
	return Result[T]{
		IsSuccess: true,
		Message:   "Successful request",
		Data:      &data,
	}
}

func Failure[T any](message string) Result[T] {
	return Result[T]{
		IsSuccess: false,
		Message:   message,
		Data:      nil,
	}
}
