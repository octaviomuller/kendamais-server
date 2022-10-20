package interfaces

import (
	"time"

	"github.com/octaviomuller/kendamais-server/internal/model"
)

type UserService interface {
	Create(email, password, name, cellphone string, cpf, cnpj *string, birthday *time.Time) error
	Login(email, password string) (*model.User, error)
	Get(id string) (*model.User, error)
	Update(id, email, name, cellphone string, cpf, cnpj *string) error
}

type UserRepository interface {
	Create(user *model.User) error
	Get(user *model.User) (*model.User, error)
	Update(user *model.User) error
}
