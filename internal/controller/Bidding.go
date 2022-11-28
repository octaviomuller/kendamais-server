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

func NewBiddingController(biddingService interfaces.BiddingService) *BiddingController {
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

func (p *BiddingController) PatchBidding(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	body := model.UpdateBidding{}

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
	dueDate := body.DueDate

	err = p.biddingService.UpdateBidding(id, title, description, minimumValue, dueDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bidding updated with success",
	})

	return
}

func (p *BiddingController) PostBid(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	body := model.MakeBid{}

	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	userId := body.UserId
	value := body.Value

	err = p.biddingService.MakeBid(id, userId, value)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bid posted with success",
	})

	return
}

func (p *BiddingController) GetBidding(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	bidding, err := p.biddingService.GetBidding(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, bidding)

	return
}

func (p *BiddingController) ListBiddings(ctx *gin.Context) {
	biddings, err := p.biddingService.ListBiddings()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, biddings)

	return
}

func (p *BiddingController) DeleteBidding(ctx *gin.Context) {
	id := ctx.Params.ByName("id")

	err := p.biddingService.DeleteBidding(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Bidding deleted with success",
	})

	return
}
