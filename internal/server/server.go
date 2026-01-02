package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Server struct {
	router *gin.Engine
	db     *mongo.Database
}

func New(db *mongo.Database) *Server {
	r := gin.Default()

	s := &Server{
		router: r,
		db:     db,
	}

	s.routes()
	return s
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) routes() {
	s.router.GET("/health", s.health)
	s.router.GET("/hello", s.hello)
}

func (s *Server) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (s *Server) hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
	})
}
