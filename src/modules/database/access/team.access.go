package access

import (
	team_core "github.com/hramov/jobhelper/src/core/team"
	"github.com/hramov/jobhelper/src/modules/database/mapper"
	"github.com/hramov/jobhelper/src/modules/database/model"
	"gorm.io/gorm"
)

type TeamAccess struct {
	Device  *model.Team
	Devices []*model.Team
	DB      *gorm.DB
}

func (ta *TeamAccess) Save(team *team_core.TeamDto) (*team_core.TeamDto, error) {
	ta.Device = nil
	tm := mapper.TeamMapper{Dto: *team}
	ta.DB.Create(tm.DtoToModel())
	// ta.DB.Raw("INSERT INTO teams (title, type, chief_id, employees) VALUES ('%s', '%s', '%')")
	return team, nil
}
