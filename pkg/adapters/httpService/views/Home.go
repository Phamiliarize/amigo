package views

import (
	"net/http"
)

// API adapter exposes REST API endpoints

type Todo struct {
	Title string
	Done  bool
}
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func (v ViewCollection) Home(w http.ResponseWriter, r *http.Request) {
	theme := v.ThemesProvider.GetTheme()
	//user := r.Context().Value("user").(dto.User)

	t := RenderTemplate(theme, "index.html")

	data := TodoPageData{
		PageTitle: "Homepage",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	w.WriteHeader(http.StatusOK)
	t.Execute(w, data)
}
