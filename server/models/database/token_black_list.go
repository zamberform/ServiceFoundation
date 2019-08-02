package database

type TokenBlackList struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Token      string    `json:"name"`
}
