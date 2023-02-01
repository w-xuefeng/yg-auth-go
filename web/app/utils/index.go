package utils

import "github.com/w-xuefeng/yg-auth-go/web/app/interfaces"

func UniJson[T interface{}](
	data T,
	success bool,
	message string,
	code int64,
) interfaces.UniformResponse[T] {
	return interfaces.UniformResponse[T]{
		Success: success, Message: message, Code: code, Data: data,
	}
}
