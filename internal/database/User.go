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

func (p *UserRepository) Create(user *model.User) error {
	tx := p.db.Create(user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
