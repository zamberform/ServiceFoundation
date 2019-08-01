package database

import "time"

type Comment struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Comment   string    `json:"comment"`
	UserId    uint      `json:"-" gorm:"column:user_id"`
	User      User      `json:"user"`
	Status    int       `json:"status"`
	ArticleId uint      `json:"article_id"`
}
