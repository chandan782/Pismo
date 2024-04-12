package services

import (
	"context"
	"net/http"
	"time"

	"github.com/chandan782/Pismo/api/helpers"
	"github.com/chandan782/Pismo/api/models"
	"github.com/chandan782/Pismo/db"
	"github.com/chandan782/Pismo/db/schemas"
	"github.com/chandan782/Pismo/response"
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
)

type Services interface {
	CreateAccount(context.Context, models.CreateAccountRequest) *response.BodyStruct
	GetAccountById(context.Context, string) *response.BodyStruct
	CreateTransaction(context.Context, models.CreateTransactionRequest) *response.BodyStruct
}

type service struct {
	DB db.DBHandlerInterface
}

func New(dbHandler db.DBHandlerInterface) (Services, error) {
	return &service{
		DB: dbHandler,
	}, nil
}

func (s *service) CreateAccount(ctx context.Context, reqBody models.CreateAccountRequest) *response.BodyStruct {
	res := &response.BodyStruct{}

	account := schemas.Account{
		ID:             shortid.MustGenerate(),
		DocumentNumber: reqBody.DocumentNumber,
	}

	err := s.DB.Create(account)
	if err != nil {
		return res.Error(http.StatusInternalServerError, "failed to create account", err, nil)
	}

	resBody := models.CreateAccountResponse{
		ID: account.ID,
	}

	return res.Success(http.StatusOK, "successfully created account", resBody)
}

func (s *service) GetAccountById(ctx context.Context, id string) *response.BodyStruct {
	res := &response.BodyStruct{}

	account := schemas.Account{}
	err := s.DB.ReadByID(&account, "id = ?", id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return res.Error(http.StatusNotFound, "account not found", err, nil)
		}
		return res.Error(http.StatusInternalServerError, "failed to get account by id", err, nil)
	}

	response := models.GetAccountByIdResponse{
		ID:             account.ID,
		DocumentNumber: account.DocumentNumber,
	}

	return res.Success(http.StatusOK, "successfully fetched account by id", response)
}

func (s *service) CreateTransaction(ctx context.Context, reqBody models.CreateTransactionRequest) *response.BodyStruct {
	res := &response.BodyStruct{}

	currentTime := time.Now().UTC().Format("2006-01-02T15:04:05.999999999Z07:00")
	trx := schemas.Transaction{
		ID:              uuid.NewString(),
		AccountID:       reqBody.AccountID,
		OperationTypeID: reqBody.OperationTypeID,
		EventDate:       currentTime,
	}

	amount := helpers.GetAmount(reqBody.Amount, reqBody.OperationTypeID)
	trx.Amount = amount

	err := s.DB.Create(trx)
	if err != nil {
		return res.Error(http.StatusInternalServerError, "failed to create transaction record in db", err, nil)
	}

	response := models.CreateTransactionResponse{
		ID: trx.ID,
	}

	return res.Success(http.StatusOK, "successfully created transaction", response)
}
