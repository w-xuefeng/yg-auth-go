package interfaces

type UniformResponse[T interface{}] struct {
	Success bool   `default:"true" json:"success"`
	Message string `default:"success" json:"message"`
	Code    int64  `default:"200" json:"code"`
	Data    T      `json:"data"`
}

type ResLogin struct {
	Token string `json:"token,omitempty"`
}

type AuthUser struct {
	Id            int64  `json:"id"`
	Stuid         string `json:"stuid"`
	Name          string `json:"name"`
	Department    string `json:"department"`
	College       string `json:"college"`
	Majorclass    string `json:"majorclass"`
	Email         string `json:"email"`
	Head          string `json:"head"`
	Ifkey         int    `json:"ifkey"`
	Online        int    `json:"online"`
	Birthday      string `json:"birthday"`
	Duty          string `json:"duty"`
	Loginip       string `json:"loginip"`
	Interviewform string `json:"interviewform"`
	Phone         string `json:"phone"`
	Photo         string `json:"photo"`
	Position      int8   `json:"position"`
	PositionName  string `json:"positionName"`
	QQ            string `json:"qq"`
	QQId          string `json:"qqid"`
	Registerdate  string `json:"registerdate"`
	Sex           string `json:"sex"`
	Signcount     int64  `json:"signcount"`
	Ulevel        int16  `json:"ulevel"`
	Utype         int16  `json:"utype"`
	WXId          string `json:"wxid"`
	Token         string `json:"token"`
}
