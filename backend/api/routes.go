package api

import (
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}))

	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(loginRepo)
	authHandler := auth.NewHandler(loginService)

	//Login
	s.Router.POST("/login", authHandler.Login)
	s.Router.POST("/register", authHandler.Register)
}
