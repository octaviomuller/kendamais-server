package interfaces

import "github.com/octaviomuller/kendamais-server/internal/model"

type BiddingService interface {
	CreateBidding(title, description string, minimumValue, bidValue float64, dueDate string, createdBy string) error
	UpdateBidding(id, title, description string, minimumValue float64, dueDate string) error
	MakeBid(id, userId string, value float64) error
	GetBidding(id string) (*model.Bidding, error)
	ListBiddings() ([]*model.Bidding, error)
	DeleteBidding(id string) error
}

type BiddingRepository interface {
	CreateBidding(bidding *model.Bidding) error
	GetBidding(bidding *model.Bidding) (*model.Bidding, error)
	ListBiddings() ([]*model.Bidding, error)
	UpdateBidding(bidding *model.Bidding) error
	DeleteBidding(bidding *model.Bidding) error
}
