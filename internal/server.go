package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer(engine *gin.Engine) *Server {
	return &Server{
		Engine: engine,
	}
}

func (server *Server) Run() {
	server.Engine.Run()
}
