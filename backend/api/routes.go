package api

import (
	"github.com/WibuSOS/SiMiddleman/auth"
	"github.com/WibuSOS/SiMiddleman/todos"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	}))

	todosRepo := todos.NewRepository(s.DB)
	todosService := todos.NewService(todosRepo)
	todosHandler := todos.NewHandler(todosService)

	s.Router.GET("/", todosHandler.GetTodos)
	s.Router.POST("/send", todosHandler.CreateTodo)
	s.Router.PATCH("/updateCheck/:task_id", todosHandler.CheckTodo)
	s.Router.DELETE("/:task_id", todosHandler.DeleteTodo)

	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	s.Router.POST("/login", authHandler.Login)
}
