package application

import (
	"github.com/takashabe/go-ddd-sample/domain"
	"github.com/takashabe/go-ddd-sample/infrastructure/persistence"
)

// GetUser returns user
func GetUser(id int) (*domain.User, error) {
	conn, err := NewDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewUserRepositoryWithRDB(conn)
	return repo.Get(id)
}

// GetUsers returns user list
func GetUsers() ([]*domain.User, error) {
	conn, err := NewDBConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	repo := persistence.NewUserRepositoryWithRDB(conn)
	return repo.GetAll()
}

// AddUser saves new user
func AddUser(name string) error {
	conn, err := NewDBConnection()
	if err != nil {
		return err
	}
	defer conn.Close()

	repo := persistence.NewUserRepositoryWithRDB(conn)
	u := &domain.User{
		Name: name,
	}
	return repo.Save(u)
}
