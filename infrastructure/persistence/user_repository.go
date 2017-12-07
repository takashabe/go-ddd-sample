package persistence

import (
	"database/sql"

	"github.com/takashabe/go-ddd-sample/domain"
	"github.com/takashabe/go-ddd-sample/domain/repository"
)

// UserRepositoryImpl Implements repository.UserRepository
type UserRepositoryImpl struct {
	Conn *sql.DB
}

// NewUserRepositoryWithRDB returns initialized UserRepositoryImpl
func NewUserRepositoryWithRDB(conn *sql.DB) repository.UserRepository {
	return &UserRepositoryImpl{Conn: conn}
}

// Get returns domain.User
func (r *UserRepositoryImpl) Get(id int) (*domain.User, error) {
	row, err := r.queryRow("select id, name from users where id=?", id)
	if err != nil {
		return nil, err
	}
	u := &domain.User{}
	err = row.Scan(&u.ID, &u.Name)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// GetAll returns list of domain.User
func (r *UserRepositoryImpl) GetAll() ([]*domain.User, error) {
	rows, err := r.query("select id, name from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	us := make([]*domain.User, 0)
	for rows.Next() {
		u := &domain.User{}
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}

// Save saves domain.User to storage
func (r *UserRepositoryImpl) Save(u *domain.User) error {
	stmt, err := r.Conn.Prepare("insert into users (name) values (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Name)
	return err
}

func (r *UserRepositoryImpl) query(q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := r.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Query(args...)
}

func (r *UserRepositoryImpl) queryRow(q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := r.Conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRow(args...), nil
}
