package team_core

type TeamProviderPort interface {
	Save(team *TeamDto) (*TeamDto, error)
}
