package controller

import (
	"api-gateway/service"
	"fmt"

	"github.com/labstack/echo/v4"
)

// MyController struct represents a controller
type UserController struct {
	Service *service.UserService
}

// NewMyController creates a new MyController instance
func NewUserController(srv *service.UserService) *UserController {
	return &UserController{Service: srv}
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *UserController) CreateUserHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		Nama     string `json:"nama" validate:"required"`
		Alamat   string `json:"alamat" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.Nama, requestPayload.Alamat)

	result, err := c.Service.CreateUser(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *UserController) DeleteUserHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		Nama   string `json:"nama" validate:"required"`
		Alamat string `json:"alamat" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.Nama, requestPayload.Alamat)

	result, err := c.Service.DeleteUser(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *UserController) LoginHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		Nama   string `json:"nama" validate:"required"`
		Alamat string `json:"alamat" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.Nama, requestPayload.Alamat)

	result, err := c.Service.Login(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *UserController) DetailUserHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		Nama   string `json:"nama" validate:"required"`
		Alamat string `json:"alamat" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.Nama, requestPayload.Alamat)

	result, err := c.Service.DetailUser(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *UserController) UpdaetUserHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		Nama   string `json:"nama" validate:"required"`
		Alamat string `json:"alamat" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.Nama, requestPayload.Alamat)

	result, err := c.Service.UpdateUser(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *UserController) ListUserHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		Nama   string `json:"nama" validate:"required"`
		Alamat string `json:"alamat" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.Nama, requestPayload.Alamat)

	result, err := c.Service.ListUser(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}
