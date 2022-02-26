package server

import (
	"github.com/gin-gonic/gin"
	"rings/domain/skywalker/handler"
)

type Server struct {
	handler *handler.Handler
}

func NewServer(handler *handler.Handler) *Server {
	return &Server{
		handler: handler,
	}
}

func (s *Server) Start() {
	engine := gin.Default()
	s.attachEndpoints(engine)
	if err := engine.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) attachEndpoints(engine *gin.Engine) {
	engine.GET("/skywalker", s.handler.HandleRequest)
}
