package team_core

type TeamEntity struct {
	Provider TeamProviderPort
}

func (tm *TeamEntity) CreateTeam(team *TeamDto) (*TeamDto, error) {
	return tm.Provider.Save(team)
}
