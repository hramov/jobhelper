package model

type Team struct {
	ID        uint    `gorm:"primaryKey"`
	Title     string  `json:"title"`
	Type      string  `json:"type"`
	Chief     *User   `json:"chief"`
	Employees []*User `json:"employees"`
}
