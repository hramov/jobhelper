package team_core

type TeamEntityPort interface {
	CreateTeam(team *TeamDto) (*TeamDto, error)
}
