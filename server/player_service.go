package server

import (
	"github.com/gin-gonic/gin"
)

type Player struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

func (s *server) GetPlayer(c *gin.Context) (*Player, error) {
	ID := c.Param("id")

	return &Player{ID: ID, Name: "Jaco Pastorius"}, nil
}
