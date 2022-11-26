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
		BidValue:     &bidValue,
		DueDate:      &dueDateParsed,
		CreatedBy:    user.Id,
	}

	err = p.biddingRepository.CreateBidding(bidding)
	if err != nil {
		return err
	}

	return nil
}

func (p *BiddingService) UpdateBidding(id, title, description string, minimumValue float64, dueDate string) error {
	if id == "" {
		return errors.New("Id required")
	}

	bidding, err := p.biddingRepository.GetBidding(&model.Bidding{Id: id})
	if err != nil {
		return err
	}

	if title != "" {
		bidding.Title = title
	}

	if description != "" {
		bidding.Description = description
	}

	if minimumValue != 0 {
		bidding.MinimumValue = minimumValue
	}

	if dueDate != "" {
		dueDateParsed, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			return err
		}

		bidding.DueDate = &dueDateParsed
	}

	err = p.biddingRepository.UpdateBidding(bidding)
	if err != nil {
		return err
	}

	return nil
}

func (p *BiddingService) MakeBid(id, userId string, value float64) error {
	if id == "" || value == 0 {
		return errors.New("Required fields missing")
	}

	user, err := p.userRepository.GetUser(&model.User{Id: userId})
	if user == nil || err != nil {
		return errors.New("User not found")
	}

	bidding, err := p.biddingRepository.GetBidding(&model.Bidding{Id: id})
	if err != nil {
		return err
	}

	if bidding.BidValue == nil && bidding.MinimumValue > value {
		return errors.New("Bid lower than minimum value")
	}

	if *bidding.BidValue >= value {
		return errors.New("Bid lower than current bid")
	}

	bidding.LastBidUser = &userId
	bidding.BidValue = &value

	err = p.biddingRepository.UpdateBidding(bidding)
	if err != nil {
		return err
	}

	return nil
}

func (p *BiddingService) GetBidding(id string) (*model.Bidding, error) {
	if id == "" {
		return nil, errors.New("Id required")
	}

	bidding, err := p.biddingRepository.GetBidding(&model.Bidding{Id: id})
	if err != nil {
		return nil, err
	}

	return bidding, nil
}

func (p *BiddingService) ListBiddings() ([]*model.Bidding, error) {
	biddings, err := p.biddingRepository.ListBiddings()
	if err != nil {
		return nil, err
	}

	return biddings, nil
}

func (p *BiddingService) DeleteBidding(id string) error {
	if id == "" {
		return errors.New("Id required")
	}

	err := p.biddingRepository.DeleteBidding(&model.Bidding{Id: id})
	if err != nil {
		return err
	}

	return nil
}
