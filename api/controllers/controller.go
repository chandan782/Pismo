package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/chandan782/Pismo/api/models"
	"github.com/chandan782/Pismo/api/services"
	"github.com/chandan782/Pismo/db"
	"github.com/chandan782/Pismo/response"
	"github.com/chandan782/Pismo/validate"
	"github.com/gofiber/fiber/v2"
)

// controllers interface defines the methods for handling API routes
type Controllers interface {
	CreateAccount(*fiber.Ctx) error
	GetAccountById(*fiber.Ctx) error
	CreateTransaction(*fiber.Ctx) error
}

type controller struct {
	service services.Services
	v       validate.Validate
}

// New creates a new instance of the controller
func New(dbHandler db.DBHandlerInterface) (Controllers, error) {
	s, err := services.New(dbHandler)
	if err != nil {
		return nil, err
	}

	v := validate.New()

	return &controller{
		service: s,
		v:       v,
	}, nil
}

// CreateAccount creates a new account
// @Summary Create a new account
// @Description Create a new account with the given details
// @Tags accounts
// @Accept json
// @Produce json
// @Param body body models.CreateAccountRequest true "Account details"
// @Success 201 {object} models.CreateAccountRequest
// @Router /accounts [post]
func (c *controller) CreateAccount(ctx *fiber.Ctx) error {
	req, res := models.CreateAccountRequest{}, &response.BodyStruct{}

	err := ctx.BodyParser(&req)
	if err != nil {
		return res.Error(http.StatusBadRequest, "failed to parse request", err, nil).Send(ctx)
	}

	err = c.v.ValidateStructs(&req)
	if err != nil {
		return res.Error(http.StatusBadRequest, "validation error", strings.ToLower(err.Error()), nil).Send(ctx)
	}

	res = c.service.CreateAccount(ctx.UserContext(), req)
	if res.Err != nil {
		return res.Error(res.StatusCode, res.Message, res.Err, nil).Send(ctx)
	}

	return res.Success(res.StatusCode, res.Message, res.Data).Send(ctx)
}

// GetAccountById retrieves account details by ID
// @Summary Get account by ID
// @Description Get account details by its ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param accountId path string true "Account ID"
// @Success 200 {object} models.GetAccountByIdResponse
// @Router /accounts/{id} [get]
func (c *controller) GetAccountById(ctx *fiber.Ctx) error {
	res := &response.BodyStruct{}
	id, err := url.PathUnescape(ctx.Params("id"))
	if err != nil || id == "" {
		return res.Error(http.StatusInternalServerError, "failed to parse param accountId value", err, nil).Send(ctx)
	}

	res = c.service.GetAccountById(ctx.UserContext(), id)
	if res.Err != nil {
		return res.Error(res.StatusCode, res.Message, res.Err, nil).Send(ctx)
	}

	return res.Success(res.StatusCode, res.Message, res.Data).Send(ctx)
}

// CreateTransaction creates a new transaction
// @Summary Create a new transaction
// @Description Create a new transaction with the given details
// @Tags transactions
// @Accept json
// @Produce json
// @Param body body models.CreateTransactionRequest true "Transaction details"
// @Success 201 {object} models.CreateTransactionRequest
// @Router /transactions [post]
func (c *controller) CreateTransaction(ctx *fiber.Ctx) error {
	req, res := models.CreateTransactionRequest{}, &response.BodyStruct{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return res.Error(http.StatusBadRequest, "failed to parse request", err, nil).Send(ctx)
	}

	err = c.v.ValidateStructs(&req)
	if err != nil {
		return res.Error(http.StatusBadRequest, "validation error", strings.ToLower(err.Error()), nil).Send(ctx)
	}

	res = c.service.CreateTransaction(ctx.UserContext(), req)
	if res.Err != nil {
		return res.Error(res.StatusCode, res.Message, res.Err, nil).Send(ctx)
	}

	return res.Success(res.StatusCode, res.Message, res.Data).Send(ctx)
}
