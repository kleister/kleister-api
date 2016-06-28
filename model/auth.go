package model

// Auth is a simple mapping for the user login form values.
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
