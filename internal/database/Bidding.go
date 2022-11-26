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

func (p *BiddingRepository) GetBidding(bidding *model.Bidding) (*model.Bidding, error) {
	result := &model.Bidding{}

	tx := p.db.Where(bidding).First(result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (p *BiddingRepository) ListBiddings() ([]*model.Bidding, error) {
	result := []*model.Bidding{}

	tx := p.db.Find(result)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (p *BiddingRepository) UpdateBidding(bidding *model.Bidding) error {
	tx := p.db.Save(bidding)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p *BiddingRepository) DeleteBidding(bidding *model.Bidding) error {
	tx := p.db.Delete(bidding)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
