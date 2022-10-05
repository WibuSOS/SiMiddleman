package api

import (
	//"github.com/WibuSOS/sinarmas/middlewares/authentication"
	//"github.com/WibuSOS/sinarmas/middlewares/authorization"

	"github.com/WibuSOS/sinarmas/auth"
	"github.com/WibuSOS/sinarmas/users"

	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}))

	// customer := []string{"consumer"}
	// admin := []string{"admin"}
	// all := []string{"customer, admin"}

	// isConsumer := authorization.Roles{AllowedRoles: customer[:]}
	// isAdmin := authorization.Roles{AllowedRoles: admin[:]}
	// isAll := authorization.Roles{AllowedRoles: all[:]}

	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	//auth
	s.Router.POST("/login", authHandler.Login)

	usersRepo := users.NewRepository(s.DB)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	s.Router.GET("/", usersHandler.GetUser)
	s.Router.POST("/register" /*authentication.Authentication, isAdmin.Authorize,*/, usersHandler.CreateUser)
	// s.Router.PATCH("/updateCheck/:task_id", usersHandler.UpdateUser)
	// s.Router.DELETE("/:task_id", usersHandler.DeleteUser)
}
