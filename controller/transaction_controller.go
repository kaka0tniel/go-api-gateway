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

type LoanPayload struct {
	IdUser string `json:"idUser" validate:"required"`
	Total  string `json:"total" validate:"required"`
	Tenor  string `json:"tenor" validate:"required"`
}

type InsPayload struct {
	IdLoan      string `json:"idLoan" validate:"required"`
	Total       string `json:"total" validate:"required"`
	TglJthTempo string `json:"tglJthTempo" validate:"required"`
}

type GolPayload struct {
	Golongan string `json:"golongan" validate:"required"`
	Total    string `json:"total" validate:"required"`
}

type LoaPayload struct {
	LoanId string `json:"idLoan" validate:"required"`
	Status string `json:"status" validate:"required"`
}

type LoatPayload struct {
	LoanId string `json:"loanId" validate:"required"`
	Total  string `json:"total" validate:"required"`
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Create Loan
// @Tags Loan
// @Accept json
// @Param user body LoanPayload true "User object in JSON format" example={"nama":"John Doe","email":"john@example.com","role":"admin","password":"secret"}
// @Produce json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /customer/create-loan [post]
func (c *TransactionController) CreateLoanHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		IdUser string `json:"idUser" validate:"required"`
		Total  string `json:"total" validate:"required"`
		Tenor  string `json:"tenor" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.IdUser, requestPayload.Total)

	result, err := c.Service.CreateLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Create Installment
// @Tags Loan
// @Accept json
// @Param user body InsPayload true "User object in JSON format" example={"nama":"John Doe","email":"john@example.com","role":"admin","password":"secret"}
// @Produce json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /customer/create-installment [post]
func (c *TransactionController) CreateInstallmentHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		IdLoan      string `json:"idLoan" validate:"required"`
		Total       string `json:"total" validate:"required"`
		TglJthTempo string `json:"tglJthTempo" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.IdLoan, requestPayload.Total)

	result, err := c.Service.CreateInstallment(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Detail Loan
// @Tags Loan
// @Accept json
// @Param id path int true "ID" Format(int64)
// @Security BearerAuth
// @Success 200 {string} result
// @Router /customer/detail-loan/{id} [get]
func (c *TransactionController) DetailLoanHandler(ctx echo.Context) error {
	// Validate path variable
	loanId := ctx.Param("id")
	if loanId == "" {
		return ctx.String(400, "Invalid Loan ID")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s"}`, loanId)

	result, err := c.Service.DetailLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Detail Max Loan
// @Tags Loan
// @Accept json
// @Param id path int true "ID" Format(int64)
// @Security BearerAuth
// @Success 200 {string} result
// @Router /admin/detail-max-loan/:id [get]
func (c *TransactionController) DetailMaxLoanHandler(ctx echo.Context) error {
	// Validate path variable
	loanId := ctx.Param("id")
	if loanId == "" {
		return ctx.String(400, "Invalid Loan ID")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s"}`, loanId)

	result, err := c.Service.DetailMaxLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary List Installment
// @Tags Loan
// @Accept json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /admin/list-installment [get]
func (c *TransactionController) ListInstallmentHandler(ctx echo.Context) error {

	result, err := c.Service.ListInstallment()
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary List Loan admin & customer
// @Tags Loan
// @Accept json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /admin/list-loan [get]
func (c *TransactionController) ListLoanHandler(ctx echo.Context) error {

	result, err := c.Service.ListLoan()
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary List Max Loan admin & customer
// @Tags Loan
// @Accept json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /admin/list-max-loan [get]
func (c *TransactionController) ListMaxLoanHandler(ctx echo.Context) error {

	result, err := c.Service.ListMaxLoan()
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Max Loan admin & customer
// @Tags Loan
// @Accept json
// @Security BearerAuth
// @Param id path int true "ID" Format(int64)
// @Param user body GolPayload true "User object in JSON format" example={"nama":"John Doe","email":"john@example.com","role":"admin","password":"secret"}
// @Success 200 {string} result
// @Router /admin/update-max-loan/:id [put]
func (c *TransactionController) UpdateMaxLoanHandler(ctx echo.Context) error {
	// Validate path variable
	// Parse path variable
	loanId := ctx.Param("id")

	// Parse request body to get payload
	var requestPayload struct {
		Golongan string `json:"golongan" validate:"required"`
		Total    string `json:"total" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	postData := fmt.Sprintf(`{"userID":"%s","title":"%s","body":"%s"}`, loanId, requestPayload.Golongan, requestPayload.Total)

	result, err := c.Service.UpdateMaxLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Update Status Loan
// @Tags Loan
// @Accept json
// @Security BearerAuth
// @Param id path int true "ID" Format(int64)
// @Param user body LoaPayload true "User object in JSON format" example={"nama":"John Doe","email":"john@example.com","role":"admin","password":"secret"}
// @Success 200 {string} result
// @Router /admin/update-status-loan/:id [patch]
func (c *TransactionController) UpdateStatusLoanHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		LoanId string `json:"idLoan" validate:"required"`
		Status string `json:"status" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.LoanId, requestPayload.Status)

	result, err := c.Service.UpdateStatusLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Customer History Installment
// @Tags Loan
// @Accept json
// @Security BearerAuth
// @Success 200 {string} result
// @Router /customer/history-installment [get]
func (c *TransactionController) HistoryInstallmentHandler(ctx echo.Context) error {

	result, err := c.Service.HistoryInstallment()
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}

// CallExternalAPIPostHandler is a controller handler for making a POST request to an external API
// @Summary Update Loan
// @Tags Loan
// @Accept json
// @Param id path int true "ID" Format(int64)
// @Param user body LoatPayload true "User object in JSON format" example={"nama":"John Doe","email":"john@example.com","role":"admin","password":"secret"}
// @Security BearerAuth
// @Success 200 {string} result
// @Router /customer/update-loan/{id} [put]
func (c *TransactionController) UpdateLoanHandler(ctx echo.Context) error {
	// Parse request body to get payload
	var requestPayload struct {
		LoanId string `json:"loanId" validate:"required"`
		Total  string `json:"total" validate:"required"`
	}

	if err := ctx.Bind(&requestPayload); err != nil {
		return ctx.String(400, "Bad Request")
	}

	// Validate payload
	if err := ctx.Validate(requestPayload); err != nil {
		return ctx.String(400, "Invalid Request Payload")
	}

	// Example POST data
	postData := fmt.Sprintf(`{"title":"%s","body":"%s","userId":1}`, requestPayload.LoanId, requestPayload.Total)

	result, err := c.Service.UpdateLoan(postData)
	if err != nil {
		return ctx.String(500, "Error calling external API with POST")
	}

	return ctx.String(200, result)
}
