package access

import (
	// team_core "github.com/hramov/jobhelper/src/core/team"
	// "github.com/hramov/jobhelper/src/modules/database/mapper"
	"github.com/hramov/jobhelper/src/modules/database/model"
	"gorm.io/gorm"
)

type TeamAccess struct {
	Device  *model.Team
	Devices []*model.Team
	DB      *gorm.DB
}
