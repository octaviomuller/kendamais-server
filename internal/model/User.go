package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        string         `gorm:"primaryKey" json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Name      string         `json:"name"`
	Cpf       *string        `json:"cpf,omitempty"`
	Cnpj      *string        `json:"cnpj,omitempty"`
	Cellphone string         `json:"cellphone"`
	Birthday  *time.Time     `json:"birthday"`
	CreatedAt *time.Time     `json:"createdAt"`
	UpdatedAt *time.Time     `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type CreateUser struct {
	Email     string
	Password  string
	Name      string
	Cpf       string
	Cnpj      string
	Cellphone string
	Birthday  string
}

type LoginUser struct {
	Email    string
	Password string
}

type UpdateUser struct {
	Email     string
	Name      string
	Cpf       string
	Cnpj      string
	Cellphone string
}
