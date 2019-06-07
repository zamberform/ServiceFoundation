package request

type User struct {
	UserId uint   `json:"id"`
	UUID   string `json:"uuid"`
	Token  string `json:"token"`
}
