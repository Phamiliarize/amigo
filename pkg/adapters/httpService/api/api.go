package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"github.com/Phamiliarize/amigo/pkg/application/port"
)

// JsonAPI adapter exposes a RESTful API that returns JSON
type JsonAPI struct {
	preferenceService port.PreferencesService
}

func NewJsonAPI(preferenceService port.PreferencesService) JsonAPI {
	return JsonAPI{
		preferenceService: preferenceService,
	}
}

// GetMe returns information about the requester
func (a JsonAPI) GetMe(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(dto.User)

	model := UserJSONOutput{
		ID:    user.ID,
		Roles: user.Roles,
	}

	responseBody, err := json.Marshal(model)
	if err != nil {
		log.Fatalf("An error occurred during JSON Marshalling: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

// PutPreferences update sthe preferences of the requester
func (a JsonAPI) PatchPreferences(w http.ResponseWriter, r *http.Request) {
	var request UserPreferenceJSONInput
	user := r.Context().Value("user").(dto.User)

	w.Header().Set("content-type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	update := dto.Preference{}

	if request.ReadingMode.Valid {
		update.ReadingMode = request.ReadingMode
		// validate the value?
	}

	if request.Theme.Valid {
		// validate the value?
		update.Theme = request.Theme
	}

	err = a.preferenceService.UpdateUserPreference(user.ID, update)
	if err != nil {
		if err.Error() == "not_found" {
			w.WriteHeader(http.StatusNotFound)
		}

		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}
