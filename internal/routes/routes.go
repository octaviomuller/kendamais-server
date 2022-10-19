package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	server "github.com/octaviomuller/kendamais-server/internal"
)

func SetupRouter(server *server.Server) {
	router := server.Engine.Group("/api/v1")
	{
		router.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "Hello, World!")
			return
		})
	}

	router.Use(cors.Default())
}
