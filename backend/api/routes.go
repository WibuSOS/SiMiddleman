package api

import (
	"github.com/WibuSOS/sinarmas/backend/middlewares/authentication"
	"github.com/WibuSOS/sinarmas/backend/middlewares/authorization"
	"github.com/WibuSOS/sinarmas/backend/middlewares/localizator"

	"github.com/WibuSOS/sinarmas/backend/controllers/auth"
	"github.com/WibuSOS/sinarmas/backend/controllers/product"
	"github.com/WibuSOS/sinarmas/backend/controllers/rooms"
	"github.com/WibuSOS/sinarmas/backend/controllers/transaction"
	"github.com/WibuSOS/sinarmas/backend/controllers/users"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() error {
	localizatorHandler, err := localizator.NewHandler()
	if err != nil {
		return err
	}

	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
	}), localizatorHandler.PassLocalizator)

	consumer := []string{"consumer"}

	consumerService := authorization.NewServiceAuthorization(s.DB, consumer)
	consumerHandler := authorization.NewHandlerAuthorization(consumerService)

	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	//auth controller (login)
	s.Router.POST("/:lang/login", authHandler.Login)

	// users controller (register)
	usersRepo := users.NewRepository(s.DB)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	s.Router.POST("/:lang/register", usersHandler.CreateUser)

	// rooms controller (create)
	roomsRepo := rooms.NewRepository(s.DB)
	roomsService := rooms.NewService(roomsRepo)
	roomsHandler := rooms.NewHandler(roomsService)

	langRoutes := s.Router.Group("/:lang")
	{
		roomRoutes := langRoutes.Group("/rooms", authentication.Authentication, consumerHandler.RoleAuthorize)
		{
			roomRoutes.POST("/", roomsHandler.CreateRoom)
			roomRoutes.GET("/:id", roomsHandler.GetAllRooms)
			roomRoutes.GET("/join/:room_id/:user_id", roomsHandler.JoinRoom)
			roomRoutes.PUT("/join/:room_id/:user_id", roomsHandler.JoinRoomPembeli)
		}
	}

	// product
	productRepo := product.NewRepository(s.DB)
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)

	// transaction
	transactionRepo := transaction.NewRepository(s.DB)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := transaction.NewHandler(transactionService)

	s.Router.PUT("/:lang/updateproduct/:id", authentication.Authentication, consumerHandler.RoleAuthorize, consumerHandler.RoomAuthorize, productHandler.UpdateProduct)
	s.Router.PUT("/:lang/updatestatus/:room_id", authentication.Authentication, consumerHandler.RoleAuthorize, consumerHandler.RoomAuthorize, transactionHandler.UpdateStatusDelivery)
	s.Router.GET("/:lang/getHarga/:room_id", authentication.Authentication, consumerHandler.RoleAuthorize, consumerHandler.RoomAuthorize, transactionHandler.GetPaymentDetails)

	return nil
}
