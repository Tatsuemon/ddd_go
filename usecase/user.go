package usecase

import (
	"context"

	"github.com/Tatsuemon/ddd_go/domain/model"
	"github.com/Tatsuemon/ddd_go/domain/repository"
	"github.com/Tatsuemon/ddd_go/infrastructure/datastore"
)

// APIのハンドラ用で使用するメソッドのインターフェース
type UserUseCase interface {
	GetUsers() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User, id string) (*model.User, error)
	DeleteUser(id string) error
}

type userUseCase struct {
	repository.UserRepository
	transaction datastore.Transaction
}

// コンストラクタ
func NewUserUseCase(r repository.UserRepository, t datastore.Transaction) UserUseCase {
	return &userUseCase{r, t}
}

func (u *userUseCase) GetUsers() ([]*model.User, error) {
	return u.UserRepository.FindAll()
}

func (u *userUseCase) GetUser(id string) (*model.User, error) {
	return u.UserRepository.FindByID(id)
}

func (u *userUseCase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	// トランザクション開始
	v, err := u.transaction.DoInTx(ctx, func(ctx context.Context) (interface{}, error) {
		return u.UserRepository.Store(ctx, user)
	})
	if err != nil {
		return nil, err
	}
	return v.(*model.User), err
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User, id string) (*model.User, error) {
	user.ID = id
	v, err := u.transaction.DoInTx(ctx, func(ctx context.Context)(interface{}, error) {
		if _, err := u.UserRepository.FindByID(user.ID); err != nil {
			return nil, err
		}
		return u.UserRepository.Update(ctx, user)
	})
	if err != nil {
		return nil, err
	}

	return v.(*model.User), err
	// return u.UserRepository.FindByID(user.ID)
}

func (u *userUseCase) DeleteUser(id string) error {
	user, err := u.UserRepository.FindByID(id)
	if err != nil {
		return err
	}
	return u.UserRepository.Delete(user)
}
