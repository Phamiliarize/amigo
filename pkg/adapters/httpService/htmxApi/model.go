package api

type UserJSONOutput struct {
	ID    string   `json:"id"`
	Roles []string `json:"roles"`
}
