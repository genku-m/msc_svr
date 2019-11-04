package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type config struct{}

type server struct {
	config *config
}

func NewConfig() *config {
	return &config{}
}

func NewServer(cfg *config) *server {
	return &server{
		config: cfg,
	}
}

func (s *server) Listen() error {
	router := gin.Default()

	router.GET("/players/:id", func(c *gin.Context) {
		res, err := s.GetPlayer(c)
		if err != nil {
			c.JSON(404, gin.H{})
			return
		}
		c.JSON(200, res)
	})
	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	return router.Run()
}
