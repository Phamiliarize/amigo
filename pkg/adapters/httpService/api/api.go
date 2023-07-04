package api

import (
	"encoding/json"
	"fmt"
	"github.com/Phamiliarize/amigo/pkg/application/dto"
	"net/http"
)

// JsonAPI adapter exposes a RESTful API that returns JSON
type JsonAPI struct {}

func NewJsonAPI() JsonAPI {
	return JsonAPI{}
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
		fmt.Printf("An error occurred during JSON Marshalling: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
