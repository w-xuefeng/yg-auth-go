package config

const (
	AuthIndex = "https://auth.youngon.work"
	AuthBase  = "https://api.youngon.work"
)

func WithBase(path string) string {
	return AuthBase + path
}

var (
	UrlLogin   = WithBase("/session")
	UrlUsers   = WithBase("/users")
	UrlRegCode = WithBase("/commonset/index/getrcode")
)
