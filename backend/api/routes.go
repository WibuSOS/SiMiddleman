package api

import (
	"github.com/WibuSOS/sinarmas/middlewares/authentication"
	"github.com/WibuSOS/sinarmas/middlewares/authorization"

	"github.com/WibuSOS/sinarmas/controllers/auth"
	"github.com/WibuSOS/sinarmas/controllers/product"
	"github.com/WibuSOS/sinarmas/controllers/rooms"
	"github.com/WibuSOS/sinarmas/controllers/users"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Accept", "Content-Type", "Authorization"},
	}))

	consumer := []string{"consumer"}
	// admin := []string{"admin"}
	// all := []string{"consumer, admin"}

	isConsumer := authorization.Roles{AllowedRoles: consumer[:]}
	// isAdmin := authorization.Roles{AllowedRoles: admin[:]}
	// isAll := authorization.Roles{AllowedRoles: all[:]}

	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	//auth
	s.Router.POST("/login", authHandler.Login)

	// users controller (register)
	usersRepo := users.NewRepository(s.DB)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	// s.Router.GET("/", usersHandler.GetUser)
	s.Router.POST("/register" /*authentication.Authentication, isAdmin.Authorize,*/, usersHandler.CreateUser)
	// s.Router.PATCH("/updateCheck/:task_id", usersHandler.UpdateUser)
	// s.Router.DELETE("/:task_id", usersHandler.DeleteUser)

	//product
	productRepo := product.NewRepository(s.DB)
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)

	// s.Router.GET("/product/:idroom", productHandler.GetSpesifikProduct)
	// s.Router.POST("/createproduct/:idroom", productHandler.CreateProduct)
	// s.Router.POST("/createproductreturnid/:idroom", productHandler.CreateProductReturnID)
	s.Router.PUT("/updateproduct/:id", authentication.Authentication, isConsumer.Authorize, productHandler.UpdateProduct)
	// s.Router.DELETE("/deleteproduct/:id", productHandler.DeleteProduct)

	// rooms controller (create)
	roomsRepo := rooms.NewRepository(s.DB)
	roomsService := rooms.NewService(roomsRepo)
	roomsHandler := rooms.NewHandler(roomsService)

	s.Router.POST("/rooms", roomsHandler.CreateRoom)
	s.Router.GET("/rooms/:id", roomsHandler.GetAllRooms)
	s.Router.GET("/joinroom/:room_id", roomsHandler.JoinRoom)
}
