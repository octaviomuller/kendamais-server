package database

import (
	"github.com/octaviomuller/kendamais-server/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (p *UserRepository) CreateUser(user *model.User) error {
	tx := p.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p *UserRepository) GetUser(user *model.User) (*model.User, error) {
	result := &model.User{}

	tx := p.db.Where(user).First(result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (p *UserRepository) UpdateUser(user *model.User) error {
	tx := p.db.Save(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
