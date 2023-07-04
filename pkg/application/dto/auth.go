package dto

// User is the key struct used to describe users in the context of AuthN/AuthZ
type User struct {
	ID    string
	Roles []string
}
