package model

import (
	"time"

	"gorm.io/gorm"
)

type Bidding struct {
	Id           string         `gorm:"primaryKey" json:"id"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	MinimumValue float64        `json:"minimumValue"`
	BidValue     float64        `json:"bidValue"`
	LastBidValue float64        `json:"lastBidValue"`
	LastBidUser  *User          `json:"lastBidUser"`
	CreatedBy    User           `json:"createdBy"`
	DueDate      *time.Time     `json:"dueDate"`
	CreatedAt    *time.Time     `json:"createdAt"`
	UpdatedAt    *time.Time     `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
}

type CreateBidding struct {
	Title        string
	Description  string
	MinimumValue float64
	BidValue     float64
	CreatedBy    string
	DueDate      *time.Time
}
