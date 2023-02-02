package interfaces

type LegacyResponse[T interface{}] struct {
	Status  bool `default:"true" json:"status"`
	ResData T    `json:"resdata,omitempty"`
  Title string `json:"title,omitempty"`
  Error string `json:"error,omitempty"`
}

type AuthBody struct {
	Stuid string `json:"stuid,omitempty"`
	Token string `json:"Token,omitempty"`
	Regcode string `json:"regcode,omitempty"`
}
