package routes

import (
	server "github.com/octaviomuller/kendamais-server/internal"
)

func SetupRouter(server *server.Server) {
	router := server.Engine.Group("/api/v1")
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
		}
	}
}
