package views

import (
	"log"
	"net/http"

	"github.com/Phamiliarize/amigo/pkg/application/dto"
)

// API adapter exposes REST API endpoints

type Todo struct {
	Title string
	Done  bool
}
type TodoPageData struct {
	Todos []Todo
}

func (v ViewCollection) Home(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(dto.User)
	theme := v.ThemeService.GetTheme(user.ID)

	data := TodoPageData{
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

	err := v.RenderTemplate(w, theme, "home.html", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
