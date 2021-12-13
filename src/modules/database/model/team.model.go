package model

type Team struct {
	Title     string  `json:"title"`
	Type      string  `json:"type"`
	Chief     *User   `json:"chief"`
	Employees []*User `json:"employees"`
}
