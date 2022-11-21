package interfaces

import "github.com/octaviomuller/kendamais-server/internal/model"

type BiddingService interface {
	CreateBidding(title, description string, minimumValue, bidValue float64, dueDate string, createdBy string) error
}

type BiddingRepository interface {
	CreateBidding(bidding *model.Bidding) error
}
