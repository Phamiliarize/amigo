package views

import (
	"fmt"
	"html/template"

	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/Phamiliarize/amigo/pkg/application/port"
)

// ViewCollection is a collection of renderable HTML Templates, effectively the UI of Amigo
type ViewCollection struct {
	ThemesProvider port.ThemesProvider
}

func NewViewCollection(themesProvider port.ThemesProvider) ViewCollection {
	return ViewCollection{
		ThemesProvider: themesProvider,
	}
}

// RenderTemplate is a helper function for rendering templates with "fallback". At default,
// if a template files doesn't exist in a theme the default theme's template should be used.
func RenderTemplate(theme dto.Theme, htmlPath string) *template.Template {
	var t *template.Template

	path := fmt.Sprintf("%s/html/%s", theme.Path, htmlPath)

	t, err := template.ParseFiles(path)
	if err != nil {
		fallbackPath := fmt.Sprintf("./themes/%s/html/%s", theme.BaseThemeDir, htmlPath)
		t = template.Must(template.ParseFiles(fallbackPath))
		return t
	}

	return t
}
