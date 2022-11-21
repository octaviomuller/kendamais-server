package service

import (
	"errors"
	"time"

	"github.com/octaviomuller/kendamais-server/internal/interfaces"
	"github.com/octaviomuller/kendamais-server/internal/model"
	uuid "github.com/satori/go.uuid"
)

type BiddingService struct {
	userRepository    interfaces.UserRepository
	biddingRepository interfaces.BiddingRepository
}

func NewBiddingService(userRepository interfaces.UserRepository, biddingRepository interfaces.BiddingRepository) *BiddingService {
	return &BiddingService{
		userRepository:    userRepository,
		biddingRepository: biddingRepository,
	}
}

func (p *BiddingService) CreateBidding(title, description string, minimumValue, bidValue float64, dueDate string, createdBy string) error {
	if title == "" || description == "" || minimumValue == 0 || bidValue == 0 || dueDate == "" || createdBy == "" {
		return errors.New("Required fields missing")
	}

	user, err := p.userRepository.GetUser(&model.User{Id: createdBy})
	if user == nil || err != nil {
		return errors.New("User not found")
	}

	dueDateParsed, err := time.Parse("2006-01-02", dueDate)
	if err != nil {
		return err
	}

	bidding := &model.Bidding{
		Id:           uuid.NewV4().String(),
		Title:        title,
		Description:  description,
		MinimumValue: minimumValue,
		BidValue:     bidValue,
		DueDate:      &dueDateParsed,
		CreatedBy:    user.Id,
	}

	err = p.biddingRepository.CreateBidding(bidding)
	if err != nil {
		return err
	}

	return nil
}
