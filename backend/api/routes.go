package api

import (
	"github.com/WibuSOS/sinarmas/login"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}))

	loginRepo := login.NewRepository(s.DB)
	loginService := login.NewService(loginRepo)
	loginHandler := login.NewHandler(loginService)

	//Login
	s.Router.POST("/login", authHandler.Login)
}
