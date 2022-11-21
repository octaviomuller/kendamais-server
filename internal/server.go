package server

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/kendamais-server/internal/controller"
	"gorm.io/gorm"
)

type Server struct {
	Engine            *gin.Engine
	DB                *gorm.DB
	UserController    controller.UserController
	BiddingController controller.BiddingController
}

func NewServer(engine *gin.Engine, db *gorm.DB, userController controller.UserController, biddingController controller.BiddingController) *Server {
	return &Server{
		Engine:            engine,
		DB:                db,
		UserController:    userController,
		BiddingController: biddingController,
	}
}

func (server *Server) Run() {
	server.Engine.Run()
}
