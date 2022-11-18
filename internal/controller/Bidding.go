package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/kendamais-server/internal/interfaces"
	"github.com/octaviomuller/kendamais-server/internal/model"
)

type BiddingController struct {
	biddingService interfaces.BiddingService
}

func NewBiddinController(biddingService interfaces.BiddingService) *BiddingController {
	return &BiddingController{
		biddingService: biddingService,
	}
}

func (p *BiddingController) PostBidding(ctx *gin.Context) {
	body := model.CreateBidding{}

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	title := body.Title
	description := body.Description
	minimumValue := body.MinimumValue
	bidValue := body.BidValue
	dueDate := body.DueDate
	createdBy := body.CreatedBy

	err = p.biddingService.CreateBidding(title, description, minimumValue, bidValue, dueDate, createdBy)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bidding created with success",
	})

	return
}
