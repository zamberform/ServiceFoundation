package database

type App struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	Name         string `json:"name"`
	Version      string `json:"version"`
	BuildCode    uint   `json:"build_code"`
	UpdateStatus int    `json:"update_status"`
	URL          string `json:"url"`
	PlatformId   int    `json:"platform_id"`
}
