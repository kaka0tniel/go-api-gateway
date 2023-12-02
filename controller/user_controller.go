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

type CreateUserRequest struct {
	Nama     string `json:"nama" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Create User
// @Description Membuat user baru
// @Tags User
// @Accept json
// @Param user body CreateUserRequest true "User object in JSON format" example={"nama":"John Doe","email":"john@example.com","role":"admin","password":"secret"}
// @Produce json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /admin/create-user [post]
func (c *UserController) CreateUserHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		Nama     string `json:"nama" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Role     string `json:"role" validate:"required"`
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
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.Nama, requestPayload.Email)

	result, err := c.Service.CreateUser(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Delete User
// @Description Menghapus user by id
// @Tags User
// @Accept json
// @Param id path int true "User ID" Format(int64)
// @Produce json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /admin/delete-user/{id} [delete]
func (c *UserController) DeleteUserHandler(ctx echo.Context) error {
	// Parse request body to get payload
	userID := ctx.Param("id")
	if userID == "" {
		return ctx.String(400, "Invalid User ID")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s"}`, userID)

	result, err := c.Service.DeleteUser(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Login
// @Description login
// @Tags Login
// @Accept json
// @Param user body LoginUser true "User object in JSON format" example={"nama":"John Doe","email":"john@example.com","role":"admin","password":"secret"}
// @Produce json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /login [post]
func (c *UserController) LoginHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		Email    string `json:"email" validate:"required"`
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
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.Email, requestPayload.Password)

	result, err := c.Service.Login(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Register User
// @Description Detail User
// @Tags User
// @Accept json
// @Param id path int true "User ID" Format(int64)
// @Produce json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /admin/detail-user/{id} [get]
func (c *UserController) DetailUserHandler(ctx echo.Context) error {

	userID := ctx.Param("id")
	if userID == "" {
		return ctx.String(400, "Invalid User ID")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s"}`, userID)

	result, err := c.Service.DetailUser(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Param id path int true "User ID" Format(int64)
// @Produce json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /admin/update-user/{id} [put]
func (c *UserController) UpdaetUserHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		Nama     string `json:"nama" validate:"required"`
		Email    string `json:"email" validate:"required"`
		Role     string `json:"role" validate:"required"`
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
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.Nama, requestPayload.Email)

	result, err := c.Service.UpdateUser(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary List User
// @Description List dari user yang terdaftar
// @Tags User
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /admin/list-user [post]
func (c *UserController) ListUserHandler(ctx echo.Context) error {

	result, err := c.Service.ListUser()
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}
