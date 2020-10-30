package usecase

import (
	"github.com/Tatsuemon/ddd_go/domain/model"
	"github.com/Tatsuemon/ddd_go/domain/repository"
)

// APIのハンドラ用で使用するメソッドのインターフェース
type UserUseCase interface {
	GetUsers() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User, id string) (*model.User, error)
	DeleteUser(id string) error
}

type userUseCase struct {
	repository.UserRepository
}

// コンストラクタ
func NewUserUseCase(r repository.UserRepository) UserUseCase {
	return &userUseCase{r}
}

func (u *userUseCase) GetUsers() ([]*model.User, error) {
	return u.UserRepository.FindAll()
}

func (u *userUseCase) GetUser(id string) (*model.User, error) {
	return u.UserRepository.FindByID(id)
}

func (u *userUseCase) CreateUser(user *model.User) (*model.User, error) {
	return u.UserRepository.Store(user)
}

func (u *userUseCase) UpdateUser(user *model.User, id string) (*model.User, error) {
	user.ID = id
	return u.UserRepository.Update(user)
}

func (u *userUseCase) DeleteUser(id string) error {
	user, err := u.UserRepository.FindByID(id)
	if err != nil {
		return err
	}
	return u.UserRepository.Delete(user)
}
