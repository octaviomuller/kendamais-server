package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	server "github.com/octaviomuller/kendamais-server/internal"
	"github.com/octaviomuller/kendamais-server/internal/config"
	"github.com/octaviomuller/kendamais-server/internal/controller"
	"github.com/octaviomuller/kendamais-server/internal/database"
	"github.com/octaviomuller/kendamais-server/internal/routes"
	"github.com/octaviomuller/kendamais-server/internal/service"
)

var connectionString string

func getEnv() {
	connectionString = os.Getenv("POSTGREE_URI")
	if connectionString == "" {
		log.Fatal("You must set your 'POSTGREE_URI' env variable.")
	}
}

func main() {
	engine := gin.Default()
	getEnv()

	db := config.ConnectDB(connectionString)

	userController := controller.NewUserController(service.NewUserService(database.NewUserRepository(db)))

	server := server.NewServer(engine, db, *userController)
	routes.SetupRouter(server)

	server.Run()
}
