package routes

import (
	server "github.com/octaviomuller/kendamais-server/internal"
)

func SetupRouter(server *server.Server) {
	router := server.Engine.Group("/api/v1")
	{
		user := router.Group("/user")
		{
			user.POST("/", server.UserController.Post)
			user.POST("/login", server.UserController.Login)
		}
	}
}
