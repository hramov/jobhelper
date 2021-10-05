package device_core

import "time"

type DeviceDto struct {
	ID          uint      `json:"id"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	InvNumber   string    `json:"inv_number"`
	Station     string    `json:"station"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	PrevCheck   time.Time `json:"prev_check"`
	NextCheck   time.Time `json:"next_check"`
}
