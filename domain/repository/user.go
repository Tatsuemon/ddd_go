package repository

import (
	"context"

	"github.com/Tatsuemon/ddd_go/domain/model"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(id string) (*model.User, error)
	FindByName(name string) (*model.User, error)
	Store(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(user *model.User) error
}
