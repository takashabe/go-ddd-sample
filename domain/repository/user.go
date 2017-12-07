package repository

import "github.com/takashabe/go-ddd-sample/domain"

// UserRepository represent repository of the user
// Expect implementation by the infrastructure layer
type UserRepository interface {
	Get(id int) (*domain.User, error)
	GetAll() ([]*domain.User, error)
	Save(*domain.User) error
}
