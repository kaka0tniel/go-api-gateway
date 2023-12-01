package main

import (
	"api-gateway/config"
	"api-gateway/controller"
	"api-gateway/service"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// Initialize configuration
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("Error initializing configuration:", err)
	}

	// Initialize Echo
	e := echo.New()

	// Set up logging to terminal
	e.Logger.SetLevel(0) // Set the logging level, 0 for DEBUG, 3 for ERROR

	// Use middleware for validation
	e.Validator = &CustomValidator{Validator: validator.New()}

	// Use middleware for logging requests and responses
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${status} ${method} ${host}${path} ${latency_human}\n",
	}))

	// Initialize service
	userService := service.NewUserService(cfg)
	transactionService := service.NewTransactionService(cfg)

	// Initialize controller
	userController := controller.NewUserController(userService)
	transactionController := controller.NewTransactionController(transactionService)

	// Define routes
	e.POST("/admin/create-user", userController.CreateUserHandler)
	e.POST("/admin/delete-user/:id", userController.DeleteUserHandler)
	e.GET("/admin/detail-loan/:id", transactionController.DetailLoanHandler)
	e.GET("/admin/detail-max-loan/:id", transactionController.DetailMaxLoanHandler)
	e.GET("/admin/detail-user", userController.DetailUserHandler)
	e.GET("/admin/list-installment", transactionController.ListInstallmentHandler)
	e.GET("/admin/list-loan", transactionController.ListLoanHandler)
	e.GET("/admin/list-max-loan", transactionController.ListMaxLoanHandler)
	e.GET("/admin/list-user", userController.ListUserHandler)
	e.PUT("/admin/update-max-loan/:id", transactionController.UpdateMaxLoanHandler)
	e.PATCH("/admin/update-status-loan/:id", transactionController.UpdateStatusLoanHandler)
	e.PUT("/admin/update-user/:id", userController.UpdaetUserHandler)
	e.POST("/customer/create-installment", transactionController.CreateInstallmentHandler)
	e.POST("/customer/create-loan", transactionController.CreateLoanHandler)
	e.GET("/customer/detail-loan/:id", transactionController.DetailLoanHandler)
	e.GET("/customer/history-installment", transactionController.HistoryInstallmentHandler)
	e.GET("/customer/list-loan", transactionController.ListLoanHandler)
	e.GET("/customer/profile", userController.DetailUserHandler)
	e.PUT("/customer/update-loan/:id", transactionController.UpdateLoanHandler)
	e.POST("/login", userController.LoginHandler)
	e.POST("/register", userController.CreateUserHandler)

	// Start the server
	log.Fatal(e.Start(":" + cfg.Port))
}

// CustomValidator is a custom validator to use with Echo
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate is a method to perform validation using the custom validator
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
