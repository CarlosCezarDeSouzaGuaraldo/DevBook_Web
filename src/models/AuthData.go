package models

// AuthData has the token and ID about the authenticated user
type AuthData struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
