package service

import (
	"errors"

	"github.com/octaviomuller/kendamais-server/internal/interfaces"
	"github.com/octaviomuller/kendamais-server/internal/model"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type BiddingService struct {
	biddingRepository interfaces.BiddingRepository
}

func NewBiddingService(biddingRepository interfaces.BiddingRepository) *BiddingService {
	return &BiddingService{
		biddingRepository: biddingRepository,
	}
}

func (p *BiddingService) CreateBidding(title, description string, minimumValue, bidValue float64, dueDate string, createdBy string) error {
	if title == "" || description == "" || minimumValue == nil || bidValue == nil || dueDate == "" || createdBy == "" {
		return errors.New("Required fields missing")
	}

	foundBidding, err := p.biddingRepository.Get(&model.Bidding{Email: email})
	if foundBidding != nil {
		return errors.New("Email unavailable")
	}

	if cpf == nil && cnpj == nil {
		return errors.New("Bidding must have cpf or cnpj")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 8)

	bidding := &model.Bidding{
		Id:        uuid.NewV4().String(),
		Email:     email,
		Password:  string(hashed),
		Name:      name,
		Cpf:       cpf,
		Cnpj:      cnpj,
		Cellphone: cellphone,
	}

	err = p.biddingRepository.Create(bidding)
	if err != nil {
		return err
	}

	return nil
}
