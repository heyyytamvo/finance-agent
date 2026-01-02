package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
	swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
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
	// Swagger endpoint
  s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
