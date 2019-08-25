package domain

import "fmt"

// User represent entity of the user
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NewUser initialize user
func NewUser(name string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name")
	}

	return &User{
		Name: name,
	}, nil
}
