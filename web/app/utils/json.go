package utils

import (
	"github.com/w-xuefeng/yg-auth-go/web/app/config"
	"github.com/w-xuefeng/yg-auth-go/web/app/interfaces"
)

func UniJson[T interface{}](
	data T,
	success bool,
	message string,
	code int64,
) interfaces.UniformResponse[T] {
	return interfaces.UniformResponse[T]{
		Success: success,
		Message: message,
		Code:    code,
		Data:    data,
	}
}

func UniJsonOk[T interface{}](
	data T,
) interfaces.UniformResponse[T] {
	return interfaces.UniformResponse[T]{
		Success: true,
		Message: config.ResMsgOk,
		Code:    config.ResCodeOk,
		Data:    data,
	}
}

func UniJsonFail[T interface{}](
	data T,
	code int64,
	message string,
) interfaces.UniformResponse[T] {
	return interfaces.UniformResponse[T]{
		Success: false,
		Message: message,
		Code:    code,
		Data:    data,
	}
}

func LegacyJsonOk[T interface{}](data T) interfaces.LegacyResponse[T] {
	return interfaces.LegacyResponse[T]{
		Status:  true,
		ResData: data,
	}
}

func LegacyJsonFail(titleAndError ...string) interfaces.LegacyResponse[any] {
	var rs = interfaces.LegacyResponse[any]{
		Status: false,
	}
	for i, item := range titleAndError {
		if i == 0 && item != "" {
			rs.Title = item
		}

  if i == 1 && item != ""{
			rs.Error = item
		}
	}
	return rs
}
