package service

import (
	"errors"
	"time"

	"github.com/octaviomuller/kendamais-server/internal/interfaces"
	"github.com/octaviomuller/kendamais-server/internal/model"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (p *UserService) Create(email string, password string, name string, cpf *string, cnpj *string, cellphone string, birthday *time.Time) error {
	if email == "" || password == "" || name == "" || cellphone == "" || birthday == nil {
		return errors.New("Required fields missing")
	}

	foundUser, err := p.userRepository.Get(&model.User{Email: email})
	if foundUser != nil {
		return errors.New("Email unavailable")
	}

	if (*birthday).After(time.Now().AddDate(-18, 0, 0)) {
		return errors.New("Users must be 18 years or older")
	}

	if cpf == nil && cnpj == nil {
		return errors.New("User must have cpf or cnpj")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	user := &model.User{
		Id:        uuid.NewV4().String(),
		Email:     email,
		Password:  string(hashed),
		Name:      name,
		Cpf:       cpf,
		Cnpj:      cnpj,
		Cellphone: cellphone,
		Birthday:  birthday,
	}

	err = p.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserService) Login(email string, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("Required fields missing")
	}

	user, err := p.userRepository.Get(&model.User{Email: email})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Wrong password")
	}

	return user, nil
}
