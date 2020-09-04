package service

import (
	"github.com/Tatsuemon/ddd_go/domain/repository"
)

type UserService interface {
	Exists(name string) (bool, error)
}

type userService struct {
	repository.UserRepository
}

// コンストラクタ
func NewUserService(r repository.UserRepository) UserService {
	return &userService{r}
}

func (u *userService) Exists(name string) (bool, error) {
	user, _ := u.UserRepository.FindByName(name)
	if user == nil {
		return true, nil
	}
	return false, nil
}
