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
			user.POST("/", server.UserController.Post)
			user.POST("/login", server.UserController.Login)
			user.GET("/:id", server.UserController.Get)
			user.PATCH("/:id", server.UserController.Patch)
		}
	}
}
