package api

import (
	"github.com/WibuSOS/sinarmas/auth"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}))

	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	//auth
	s.Router.POST("/login", authHandler.Login)
	//Register
	// s.Router.POST("/register", authHandler.Register)
}
