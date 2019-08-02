package database

import "time"

type Article struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	TagId       uint      `json:"tag_id" gorm:"column:tag_id"`
	Tag         Tag       `json:"-"`
	Status      int       `json:"status"`
	ContentDesc string    `json:"content_desc"`
	CommentFlg  bool      `json:"comment_flg"`
}
