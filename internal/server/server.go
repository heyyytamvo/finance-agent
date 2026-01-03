package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"
  "my-finance-app/internal/services/spending"
)

type Server struct {
	router          *gin.Engine
  SpendingService *spending.Service
}

func New(spService *spending.Service) *Server {
	r := gin.Default()

  s := &Server{
    router:          r,
    SpendingService: spService,
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
	s.router.POST("/spendings", s.createSpending)
	s.router.GET("/spendings", s.getAllSpendings)
// 	s.router.GET("/spendings/filter", s.getSpendingsByCategory)
//   s.router.GET("/spendings/total", s.getCostByCategory)
  s.router.GET("spendings/total", s.getCostByCategoryRange)
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

// CreateSpending handles POST /spendings
func (s *Server) createSpending(c *gin.Context) {
	var sp spending.Spending
	if err := c.ShouldBindJSON(&sp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	created, err := s.SpendingService.Create(c.Request.Context(), sp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// getAllSpendings handles GET /spendings
func (s *Server) getAllSpendings(c *gin.Context) {
	spendings, err := s.SpendingService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, spendings)
}

// // getAllSpendings handles GET /spending
// func (s *Server) getSpendingsByCategory(c *gin.Context) {
// 	category := c.Query("type") // GET /spendings?type=groceries
// 	if category == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "type query parameter required"})
// 		return
// 	}
//
// 	spendings, err := s.SpendingService.GetByCategory(c.Request.Context(), category)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	c.JSON(http.StatusOK, spendings)
// }


// getCostByCategory handles GET /spending/total
func (s *Server) getCostByCategory(c *gin.Context) {
	category := c.Query("type") // optional

	total, err := s.SpendingService.GetCostByCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"type":  category, // empty string if not specified
		"total": total,
	})
}

// getCostByCategoryRange handles GET /spending/total
func (s *Server) getCostByCategoryRange(c *gin.Context) {
	category := c.Query("type")

	var from, to *time.Time
	fromStr := c.Query("from")
	toStr := c.Query("to")

	if fromStr != "" {
		t, err := time.Parse(time.RFC3339, fromStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid from date"})
			return
		}
		from = &t
	}
	if toStr != "" {
		t, err := time.Parse(time.RFC3339, toStr)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid to date"})
			return
		}
		to = &t
	}

	total, err := s.SpendingService.GetCostByCategoryAndRange(c.Request.Context(), category, from, to)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"type":  category,
		"from":  fromStr,
		"to":    toStr,
		"total": total,
	})
}
