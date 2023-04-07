package models

import "time"

// User is a user from social media
type User struct {
	ID           uint64         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Nick         string         `json:"nick,omitempty"`
	Email        string         `json:"email,omitempty"`
	CreatedAt    time.Time      `json:"createdAt,omitempty"`
	Followers    []User         `json:"followers,omitempty"`
	Following    []User         `json:"following,omitempty"`
	Publications []Publications `json:"publications,omitempty"`
}
