package views

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/Phamiliarize/amigo/pkg/application/port"
)

type TemplateData struct {
	ReadingMode bool
	Metadata    dto.ThemeMetadata
	ViewData    any
}

// ViewCollection is a collection of renderable HTML Templates, effectively the UI of Amigo
type ViewCollection struct {
	ThemeService port.ThemeService
}

func NewViewCollection(ThemeService port.ThemeService) ViewCollection {
	return ViewCollection{
		ThemeService: ThemeService,
	}
}

// RenderTemplate is a helper function for rendering templates with "fallback". At default,
// if a template files doesn't exist in a given theme- fallback to the base theme's template.
func (v ViewCollection) RenderTemplate(w http.ResponseWriter, theme dto.Theme, htmlPath string, data any) error {
	var t *template.Template

	fullData := TemplateData{
		ReadingMode: theme.ReadingMode,
		Metadata:    v.ThemeService.GetCachedThemeMetadata().ThemeMetadata,
		ViewData:    data,
	}

	layoutTemplate := fmt.Sprintf("%s/html/layout.html", theme.Path)
	headTemplate := fmt.Sprintf("%s/html/head.html", theme.Path)
	footerTemplate := fmt.Sprintf("%s/html/footer.html", theme.Path)
	path := fmt.Sprintf("%s/html/%s", theme.Path, htmlPath)

	templates := []string{
		layoutTemplate,
		headTemplate,
		footerTemplate,
		path,
	}

	t, err := template.ParseFiles(templates...)
	if err != nil {
		fallbackPath := fmt.Sprintf("./themes/%s/html/%s", theme.BaseThemeDir, htmlPath)
		t, err = template.ParseFiles(fallbackPath)
		if err != nil {
			return err
		}
	}

	err = t.ExecuteTemplate(w, "layout", fullData)
	if err != nil {
		return err
	}

	return nil
}
