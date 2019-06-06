package request

type CommonReq struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	BuildNo  string `json:"build_no"`
	Platform int    `json:"platform"`
	User     User   `json:"user"`
}
