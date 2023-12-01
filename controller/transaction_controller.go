package controller

import (
	"api-gateway/service"
	"fmt"

	"github.com/labstack/echo/v4"
)

// MyController struct represents a controller
type TransactionController struct {
	Service *service.TransactionService
}

// NewMyController creates a new MyController instance
func NewTransactionController(srv *service.TransactionService) *TransactionController {
	return &TransactionController{Service: srv}
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) CreateLoanHandler(ctx echo.Context) error {
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

	result, err := c.Service.CreateInstallment(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) CreateInstallmentHandler(ctx echo.Context) error {
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

	result, err := c.Service.CreateLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) DetailLoanHandler(ctx echo.Context) error {
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

	result, err := c.Service.DetailLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) DetailMaxLoanHandler(ctx echo.Context) error {
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

	result, err := c.Service.DetailMaxLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) ListInstallmentHandler(ctx echo.Context) error {
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

	result, err := c.Service.ListInstallment(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) ListLoanHandler(ctx echo.Context) error {
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

	result, err := c.Service.ListLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) ListMaxLoanHandler(ctx echo.Context) error {
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

	result, err := c.Service.ListMaxLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) UpdateMaxLoanHandler(ctx echo.Context) error {
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

	result, err := c.Service.UpdateMaxLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) UpdateStatusLoanHandler(ctx echo.Context) error {
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

	result, err := c.Service.UpdateStatusLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) HistoryInstallmentHandler(ctx echo.Context) error {
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

	result, err := c.Service.HistoryInstallment(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
func (c *TransactionController) UpdateLoanHandler(ctx echo.Context) error {
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

	result, err := c.Service.UpdateLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}
