package main

import (
	"github.com/gin-gonic/gin"
	server "github.com/octaviomuller/kendamais-server/internal"
	"github.com/octaviomuller/kendamais-server/internal/routes"
)

func main() {
	engine := gin.Default()

	server := server.NewServer(engine)
	routes.SetupRouter(server)

	server.Run()
}
