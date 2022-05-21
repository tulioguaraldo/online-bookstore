package server

import (
	"log"

	"github.com/gin-gonic/gin"

	env "github.com/tulioguaraldo/online-bookstore/config/environment"
	"github.com/tulioguaraldo/online-bookstore/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   env.Port,
		server: gin.Default(),
	}
}

func (s *Server) RunServer() {
	router := routes.GetRoutes()
	log.Fatal(router.Run(":" + s.port))
}
