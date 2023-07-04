package port

import (
	"github.com/Phamiliarize/amigo/pkg/application/dto"
)

type ThemesProvider interface {
	GetTheme() dto.Theme
	GetThemes() []dto.Theme
}
