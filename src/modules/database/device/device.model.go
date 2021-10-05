package device_db

import "time"

type Device struct {
	ID          uint16    `gorm:"primaryKey"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	InvNumber   string    `json:"inv_number"`
	Station     string    `json:"station"`
	Location    string    `json:"location"`
	PrevCheck   time.Time `json:"prev_check"`
	NextCheck   time.Time `json:"next_check"`
}
