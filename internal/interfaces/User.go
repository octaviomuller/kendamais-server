package interfaces

import (
	"time"

	"github.com/octaviomuller/kendamais-server/internal/model"
)

type UserService interface {
	Create(email string, password string, name string, cpf *string, cnpj *string, cellphone string, birthday *time.Time) error
}

type UserRepository interface {
	Create(user *model.User) error
}
