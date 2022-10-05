package api

import (
	"github.com/WibuSOS/sinarmas/auth"
	"github.com/WibuSOS/sinarmas/product"
	"github.com/WibuSOS/sinarmas/users"
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

	usersRepo := users.NewRepository(s.DB)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	// s.Router.GET("/", usersHandler.GetUser)
	s.Router.POST("/register", usersHandler.CreateUser)
	// s.Router.PATCH("/updateCheck/:task_id", usersHandler.UpdateUser)
	// s.Router.DELETE("/:task_id", usersHandler.DeleteUser)

	//product
	productRepo := product.NewRepository(s.DB)
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)

	s.Router.GET("/product/:idroom", productHandler.GetSpesifikProduct)
	s.Router.POST("/createproduct/:idroom", productHandler.CreateProduct)
	// s.Router.POST("/createproductreturnid/:idroom", productHandler.CreateProductReturnID)
	s.Router.PUT("/updateproduct/:id", productHandler.UpdateProduct)
	s.Router.DELETE("/deleteproduct/:id", productHandler.DeleteProduct)
}
