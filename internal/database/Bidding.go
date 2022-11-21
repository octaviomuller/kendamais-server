package database

import (
	"github.com/octaviomuller/kendamais-server/internal/model"
	"gorm.io/gorm"
)

type BiddingRepository struct {
	db *gorm.DB
}

func NewBiddingRepository(db *gorm.DB) *BiddingRepository {
	return &BiddingRepository{
		db: db,
	}
}

func (p *BiddingRepository) CreateBidding(bidding *model.Bidding) error {
	tx := p.db.Create(bidding)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
