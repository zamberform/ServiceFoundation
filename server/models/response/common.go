package response

type CommonRes struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	User User   `json:"user"`
}
