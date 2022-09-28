package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Engine *gin.Engine
	DB     *gorm.DB
}

func NewServer(engine *gin.Engine, db *gorm.DB) *Server {
	return &Server{
		Engine: engine,
		DB:     db,
	}
}

func (server *Server) Run() {
	server.Engine.Run()
}
