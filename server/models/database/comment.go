package database

import "time"

type Comment struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Comment   string     `json:"comment"`
	User      User       `json:"user"`
	Status    int        `json:"status"`
	ArticleId uint       `json:"articleId"`
}
