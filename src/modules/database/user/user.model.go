package user_db

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	LastName  string    `json:"last_name"`
	Name      string    `json:"name"`
	Position  string    `json:"position"`
	Role      string    `json:"role"`
	Station   string    `json:"station"`
	ChatID    int64     `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
}
