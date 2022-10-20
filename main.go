package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	server "github.com/octaviomuller/kendamais-server/internal"
	"github.com/octaviomuller/kendamais-server/internal/config"
	"github.com/octaviomuller/kendamais-server/internal/controller"
	"github.com/octaviomuller/kendamais-server/internal/database"
	"github.com/octaviomuller/kendamais-server/internal/routes"
	"github.com/octaviomuller/kendamais-server/internal/service"
)

var connectionString string

func getEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}

	connectionString = os.Getenv("POSTGREE_URI")
	if connectionString == "" {
		log.Fatal("You must set your 'POSTGREE_URI' env variable.")
	}
}

func main() {
	engine := gin.Default()
	engine.Use(cors.Default())

	getEnv()

	db := config.ConnectDB(connectionString)

	userController := controller.NewUserController(service.NewUserService(database.NewUserRepository(db)))

	server := server.NewServer(engine, db, *userController)
	routes.SetupRouter(server)

	server.Run()
}
