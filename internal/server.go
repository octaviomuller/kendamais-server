package server

import (
	"github.com/gin-gonic/gin"
	"github.com/octaviomuller/kendamais-server/internal/controller"
	"gorm.io/gorm"
)

type Server struct {
	Engine         *gin.Engine
	DB             *gorm.DB
	UserController controller.UserController
}

func NewServer(engine *gin.Engine, db *gorm.DB, userController controller.UserController) *Server {
	return &Server{
		Engine:         engine,
		DB:             db,
		UserController: userController,
	}
}

func (server *Server) Run() {
	server.Engine.Run()
}
