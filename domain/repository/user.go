package repository

import (
	"context"

	"github.com/takashabe/go-ddd-sample/domain"
)

//go:generate mockgen -package $GOPACKAGE -source $GOFILE -destination mock_$GOFILE

// UserRepository represent repository of the user
// Expect implementation by the infrastructure layer
type UserRepository interface {
	Get(ctx context.Context, id int) (*domain.User, error)
	GetAll(ctx context.Context) ([]*domain.User, error)
	Save(ctx context.Context, user *domain.User) error
}
