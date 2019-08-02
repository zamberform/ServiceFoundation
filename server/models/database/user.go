package database

import "time"

type User struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `sql:"index" json:"deleted_at"`
	ActivatedAt *time.Time `json:"activated_at"`
	Name        string     `json:"name"`
	UUID        string     `json:"uuid"`
	Pass        string     `json:"password"`
	Email       string     `json:"email"`
	Introduce   string     `json:"introduce"`
	AvatarURL   string     `json:"avatarURL"`
	Status      int        `json:"status"`
	PlatformId  int        `json:"platform_id"`
}
