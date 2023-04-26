package users

import "errors"

var (
	ErrNotFound      = errors.New("user not found in the given repository")
	ErrAlreadyExists = errors.New("user already exists in the given repository")
)

type Repository interface {
	GetByID(id int) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
}
