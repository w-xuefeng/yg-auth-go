package interfaces

type LegacyResponse[T interface{}] struct {
	Status  bool `default:"true" json:"status"`
	ResData T    `json:"resdata"`
}

type AuthBody struct {
	Stuid string `json:"stuid"`
	Token string `json:"Token"`
	Regcode string `json:"regcode"`
}
