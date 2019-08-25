package persistence

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // driver
	"github.com/takashabe/go-ddd-sample/domain"
	"github.com/takashabe/go-ddd-sample/domain/repository"
)

// userRepository Implements repository.UserRepository
type userRepository struct {
	conn *sql.DB
}

// NewUserRepository returns initialized UserRepositoryImpl
func NewUserRepository(conn *sql.DB) repository.UserRepository {
	return &userRepository{conn: conn}
}

// Get returns domain.User
func (r *userRepository) Get(ctx context.Context, id int) (*domain.User, error) {
	row, err := r.queryRow(ctx, "select id, name from users where id=?", id)
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
func (r *userRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	rows, err := r.query(ctx, "select id, name from users")
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
func (r *userRepository) Save(ctx context.Context, u *domain.User) error {
	stmt, err := r.conn.Prepare("insert into users (name) values (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, u.Name)
	return err
}

func (r *userRepository) query(ctx context.Context, q string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryContext(ctx, args...)
}

func (r *userRepository) queryRow(ctx context.Context, q string, args ...interface{}) (*sql.Row, error) {
	stmt, err := r.conn.Prepare(q)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...), nil
}
