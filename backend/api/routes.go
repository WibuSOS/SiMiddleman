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
	// localization middleware
	localizatorHandler, err := localizator.NewHandler()
	if err != nil {
		return err
	}

	// consumer authorization middleware
	consumer := []string{"consumer"}
	consumerService := authorization.NewServiceAuthorization(s.DB, consumer)
	consumerHandler := authorization.NewHandlerAuthorization(consumerService)

	// auth controller (login)
	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	// users controller (register)
	usersRepo := users.NewRepository(s.DB)
	usersService := users.NewService(usersRepo)
	usersHandler := users.NewHandler(usersService)

	// rooms controller
	roomsRepo := rooms.NewRepository(s.DB)
	roomsService := rooms.NewService(roomsRepo)
	roomsHandler := rooms.NewHandler(roomsService)

	// product controller
	productRepo := product.NewRepository(s.DB)
	productService := product.NewService(productRepo)
	productHandler := product.NewHandler(productService)

	// transaction controller
	transactionRepo := transaction.NewRepository(s.DB)
	transactionService := transaction.NewService(transactionRepo)
	transactionHandler := transaction.NewHandler(transactionService)

	langRoutes := s.Router.Group("/:lang")
	{
		langRoutes.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowHeaders: []string{"Origin", "Accept", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		}), localizatorHandler.PassLocalizator)

		// auth controller (login)
		langRoutes.POST("/login", authHandler.Login)

		// users controller (register)
		langRoutes.POST("/register", usersHandler.CreateUser)

		roomRoutes := langRoutes.Group("/rooms", authentication.Authentication, consumerHandler.RoleAuthorize)
		{
			roomRoutes.POST("/", roomsHandler.CreateRoom)
			roomRoutes.GET("/:id", roomsHandler.GetAllRooms)
			roomRoutes.GET("/join/:room_id/:user_id", roomsHandler.JoinRoom)
			roomRoutes.PUT("/join/:room_id/:user_id", roomsHandler.JoinRoomPembeli)

			detailRoutes := roomRoutes.Group("/details")
			{
				detailRoutes.PUT("/updateproduct/:product_id", productHandler.UpdateProduct)
				detailRoutes.PUT("/updatestatus/:room_id", consumerHandler.RoomAuthorize, consumerHandler.RoomAuthorize, transactionHandler.UpdateStatusDelivery)
				detailRoutes.GET("/getHarga/:room_id", consumerHandler.RoomAuthorize, transactionHandler.GetPaymentDetails)
			}
		}
	}

	return nil
}
