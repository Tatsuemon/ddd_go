package repository

import "github.com/Tatsuemon/ddd_go/domain/model"

type UserRepository interface {
	FindByID(id string) (*model.User, error)
	FindByName(name string) (*model.User, error)
	Store(user *model.User) error
	Update(user *model.User) error
	Delete(user *model.User) error
}
