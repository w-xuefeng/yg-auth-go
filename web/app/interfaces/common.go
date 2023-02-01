package interfaces

type UniformResponse[T interface{}] struct {
	Success bool   `default:"true" json:"success"`
	Message string `default:"success" json:"message"`
	Code    int64  `default:"200" json:"code"`
	Data    T      `json:"data"`
}
