// ドメインモデル(値オブジェクトとかは作ってない)
package model

import (
	"fmt"
	"unicode/utf8"

	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

// 作ったけど使わなそう
// TODO(Tatsuemon): バリデーションをいれるべき
func NewUser(name string, email string) (*User, error) {
	if name == "" {
		return nil, fmt.Errorf("name is required.")
	}
	if utf8.RuneCountInString(name) < 3 {
		return nil, fmt.Errorf("name must be at least 3 characters.")
	}
	if email == "" {
		return nil, fmt.Errorf("email is required.")
	}
	return &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}, nil
}
