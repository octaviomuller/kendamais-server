package routes

import (
	"github.com/gin-contrib/cors"
	server "github.com/octaviomuller/kendamais-server/internal"
)

func SetupRouter(server *server.Server) {
	router := server.Engine.Group("/api/v1")
	router.Use(cors.Default())
	{
		user := router.Group("/user")
		{
			user.POST("/", server.UserController.PostUser)
			user.POST("/login", server.UserController.Login)
			user.GET("/:id", server.UserController.GetUser)
			user.PATCH("/:id", server.UserController.PatchUser)
		}

		bidding := router.Group("/bidding")
		{
			bidding.POST("/", server.BiddingController.PostBidding)
			bidding.PATCH("/:id", server.BiddingController.PatchBidding)
			bidding.POST("/:id", server.BiddingController.PostBid)
			bidding.GET("/:id", server.BiddingController.GetBidding)
			bidding.GET("/", server.BiddingController.ListBiddings)
			bidding.DELETE("/:id", server.BiddingController.DeleteBidding)
		}
	}
}
