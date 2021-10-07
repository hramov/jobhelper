package user_core

import "time"

type UserDto struct {
	ID        uint      `gorm:"primaryKey"`
	LastName  string    `json:"last_name"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Position  string    `json:"position"`
	Station   string    `json:"station"`
	ChatID    int64     `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
}
