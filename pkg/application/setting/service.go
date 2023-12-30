package setting

import (
	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/Phamiliarize/amigo/pkg/application/port"
)

type Setting struct {
	db port.Database
}

func NewSettingService(db port.Database) Setting {
	return Setting{
		db: db,
	}
}

func (s Setting) GetGeneralSetting() (*dto.GeneralSetting, error) {
	settings, err := s.db.GetGeneralSetting()
	if err != nil {
		return nil, err
	}

	return settings, nil
}
