package usecase

import (
	"github.com/Tatsuemon/ddd_go/domain/model"
	"github.com/Tatsuemon/ddd_go/unit_of_work"
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
	Uow unit_of_work.IunitOfWork
}

// コンストラクタ
func NewUserUseCase(uow unit_of_work.IunitOfWork) UserUseCase {
	return &userUseCase{
		Uow: uow,
	}
}

func (u *userUseCase) GetUsers() ([]*model.User, error) {
	return u.Uow.GetUserRepository().FindAll()
}

func (u *userUseCase) GetUser(id string) (*model.User, error) {
	return u.Uow.GetUserRepository().FindByID(id)
}

func (u *userUseCase) CreateUser(user *model.User) (*model.User, error) {
	user, err := u.Uow.GetUserRepository().Store(user)

	if err != nil {
		u.Uow.Rollback()
		return nil, err
	}

	u.Uow.Commit()
	return user, nil
}

func (u *userUseCase) UpdateUser(user *model.User, id string) (*model.User, error) {
	user.ID = id
	user, err := u.Uow.GetUserRepository().Update(user)
	if err != nil {
		u.Uow.Rollback()
		return nil, err
	}

	u.Uow.Commit()
	return user, nil
}

func (u *userUseCase) DeleteUser(id string) error {
	user, err := u.Uow.GetUserRepository().FindByID(id)
	if err != nil {
		return err
	}
	err = u.Uow.GetUserRepository().Delete(user)
	if err != nil {
		u.Uow.Rollback()
		return err
	}

	u.Uow.Commit()
	return nil
}
