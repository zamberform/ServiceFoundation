package model

type App struct {
	Name         string `json:"name"`
	Version      string `json:"version"`
	UpdateStatus int    `json:"update_status"`
	AppURL       string `json:"app_url"`
	AppPlatform  int    `json:"platform"`
}
