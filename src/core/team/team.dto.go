package team_core

import user_core "github.com/hramov/jobhelper/src/core/user"

type TeamDto struct {
	Title     string               `json:"title"`
	Type      string               `json:"type"`
	Chief     *user_core.UserDto   `json:"chief"`
	Employees []*user_core.UserDto `json:"employees"`
}
