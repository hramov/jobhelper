package mapper

import (
	team_core "github.com/hramov/jobhelper/src/core/team"
	"github.com/hramov/jobhelper/src/modules/database/model"
)

type TeamMapper struct {
	Dto   team_core.TeamDto
	Model model.Team
}

func (tm *TeamMapper) DtoToModel() *model.Team {
	team := tm.Model
	team.Title = tm.Dto.Title
	team.Type = tm.Dto.Type

	um := UserMapper{Dto: *tm.Dto.Chief}
	team.Chief = um.DtoToModel()

	for _, employee := range tm.Dto.Employees {
		um := UserMapper{Dto: *employee}
		team.Employees = append(team.Employees, um.DtoToModel())
	}

	return &team
}

func (tm *TeamMapper) ModelToDto() *team_core.TeamDto {
	team := tm.Dto
	team.Title = tm.Model.Title
	team.Type = tm.Model.Type

	um := UserMapper{Model: *tm.Model.Chief}
	team.Chief = um.ModelToDto()

	for _, employee := range tm.Model.Employees {
		um := UserMapper{Model: *employee}
		team.Employees = append(team.Employees, um.ModelToDto())
	}

	return &team
}
