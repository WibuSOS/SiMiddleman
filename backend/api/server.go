package api

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB:     db,
	}

	return s
}

func (s *server) RunServer() {
	port := os.Getenv("PORT")

	if err := s.SetupRouter(); err != nil {
		log.Panicln(err.Error())
	}

	if err := s.Router.Run(":" + port); err != nil {
		log.Panicln(err.Error())
	}
}
