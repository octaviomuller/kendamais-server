package service

import (
	"errors"

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

func (p *UserService) CreateUser(email, password, name, cellphone string, cpf, cnpj *string) error {
	if email == "" || password == "" || name == "" || cellphone == "" {
		return errors.New("Required fields missing")
	}

	foundUser, err := p.userRepository.GetUser(&model.User{Email: email})
	if foundUser != nil {
		return errors.New("Email unavailable")
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
	}

	err = p.userRepository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (p *UserService) Login(email string, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("Required fields missing")
	}

	user, err := p.userRepository.GetUser(&model.User{Email: email})
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Wrong password")
	}

	return user, nil
}

func (p *UserService) GetUser(id string) (*model.User, error) {
	if id == "" {
		return nil, errors.New("Id not specified")
	}

	user, err := p.userRepository.GetUser(&model.User{Id: id})
	if err != nil || user == nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (p *UserService) UpdateUser(id, email, name, cellphone string, cpf, cnpj *string) error {
	if id == "" {
		return errors.New("Id not specified")
	}

	user, err := p.userRepository.GetUser(&model.User{Id: id})
	if err != nil || user == nil {
		return errors.New("User not found")
	}

	if email != "" {
		user.Email = email
	}
	if name != "" {
		user.Name = name
	}
	if cellphone != "" {
		user.Cellphone = cellphone
	}
	if cpf != nil {
		user.Cpf = cpf
	}
	if cnpj != nil {
		user.Cnpj = cnpj
	}

	err = p.userRepository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}
