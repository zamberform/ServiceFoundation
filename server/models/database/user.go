package database

import "time"

type User struct {
	ID          uint       `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `sql:"index" json:"deletedAt"`
	ActivatedAt *time.Time `json:"activatedAt"`
	Name        string     `json:"name"`
	UUID        string     `json:"uuid"`
	Pass        string     `json:"-"`
	Email       string     `json:"-"`
	Introduce   string     `json:"introduce"`
	AvatarURL   string     `json:"avatarURL"`
	Status      int        `json:"status"`
	App         App        `json:"app"`
}
