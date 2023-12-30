package preferences

import (
	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/Phamiliarize/amigo/pkg/application/port"
)

type Preferences struct {
	db port.Database
}

func NewPreferencesService(db port.Database) Preferences {
	return Preferences{
		db: db,
	}
}

func (p Preferences) GetUserPreference(userID string) (*dto.Preference, error) {
	preference, err := p.db.GetUserPreference(userID)
	if err != nil {
		return nil, err
	}

	return preference, nil
}
