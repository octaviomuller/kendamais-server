package interfaces

import (
	"time"

	"github.com/octaviomuller/kendamais-server/internal/model"
)

type UserService interface {
	Create(email string, password string, name string, cpf *string, cnpj *string, cellphone string, birthday *time.Time) error
	Login(email string, password string) (*model.User, error)
}

type UserRepository interface {
	Create(user *model.User) error
	Get(user *model.User) (*model.User, error)
}
