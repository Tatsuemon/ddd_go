package repository

import "github.com/Tatsuemon/ddd_go/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindByID(id string) (*model.User, error)
	FindByName(name string) (*model.User, error)
	Store(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(user *model.User) error
}
