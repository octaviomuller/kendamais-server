package interfaces

import (
	"github.com/octaviomuller/kendamais-server/internal/model"
)

type UserService interface {
	CreateUser(email, password, name, cellphone string, cpf, cnpj *string) error
	Login(email, password string) (*model.User, error)
	GetUser(id string) (*model.User, error)
	UpdateUser(id, email, name, cellphone string, cpf, cnpj *string) error
}

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUser(user *model.User) (*model.User, error)
	UpdateUser(user *model.User) error
}
