package services_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/chandan782/Pismo/api/models"
	"github.com/chandan782/Pismo/api/services"
	"github.com/chandan782/Pismo/db/schemas"
	"github.com/chandan782/Pismo/internals/mocks"
	"github.com/chandan782/Pismo/response"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateAccount_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBHandlerInterface(ctrl)

	s, err := services.New(mockDB)
	if err != nil {
		t.Fatalf("error creating service: %v", err)
	}

	ctx := context.TODO()
	documentNumber := "1234567890"

	// set up expectations on the mock
	mockDB.EXPECT().Create(gomock.Any()).Return(nil)

	req := models.CreateAccountRequest{
		DocumentNumber: documentNumber,
	}

	res := s.CreateAccount(ctx, req)
	expectedResponse := response.BodyStruct{
		StatusCode: http.StatusOK,
	}

	assert.Equal(t, expectedResponse.StatusCode, res.StatusCode)
}

func TestCreateAccount_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBHandlerInterface(ctrl)

	s, err := services.New(mockDB)
	if err != nil {
		t.Fatalf("error creating service: %v", err)
	}

	ctx := context.TODO()
	documentNumber := "1234567890"

	// mocking db error
	expectedErr := errors.New("failed to create account")
	mockDB.EXPECT().Create(gomock.Any()).Return(expectedErr)

	req := models.CreateAccountRequest{
		DocumentNumber: documentNumber,
	}

	res := s.CreateAccount(ctx, req)
	expectedResponse := response.BodyStruct{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to create account",
	}

	assert.Equal(t, expectedResponse.StatusCode, res.StatusCode)
	assert.Equal(t, expectedResponse.Message, res.Message)
}

func TestGetAccountById_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// create mock DB handler
	mockDB := mocks.NewMockDBHandlerInterface(ctrl)

	// initialize your service with the mock DB handler
	s, err := services.New(mockDB)
	if err != nil {
		t.Fatalf("error creating service: %v", err)
	}

	accountId := "accountid"

	account := schemas.Account{
		ID:             accountId,
		DocumentNumber: "9876543210",
	}

	// expect the DB ReadByID method to be called with the correct arguments
	mockDB.EXPECT().ReadByID(gomock.Any(), "id = ?", accountId).DoAndReturn(func(_, _ interface{}, _ ...interface{}) (schemas.Account, error) {
		return account, nil
	})

	res := s.GetAccountById(context.TODO(), accountId)

	expectedResponse := response.BodyStruct{
		StatusCode: http.StatusOK,
		Message:    "successfully fetched account by id",
		Data: models.GetAccountByIdResponse{
			ID:             account.ID,
			DocumentNumber: account.DocumentNumber,
		},
	}

	assert.Equal(t, expectedResponse.StatusCode, res.StatusCode)
	assert.Equal(t, expectedResponse.Message, res.Message)
}

func TestGetAccountById_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBHandlerInterface(ctrl)

	s, err := services.New(mockDB)
	if err != nil {
		t.Fatalf("error creating service: %v", err)
	}

	ctx := context.TODO()
	accountId := "accountid"

	// expected error by DB ReadByID method
	expectedErr := errors.New("failed to get account by id")
	mockDB.EXPECT().ReadByID(gomock.Any(), "id = ?", accountId).Return(expectedErr)

	req := accountId

	res := s.GetAccountById(ctx, req)
	expectedResponse := response.BodyStruct{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to get account by id",
	}

	assert.Equal(t, expectedResponse.StatusCode, res.StatusCode)
	assert.Equal(t, expectedResponse.Message, res.Message)
}

func TestGetAccountById_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBHandlerInterface(ctrl)

	s, err := services.New(mockDB)
	if err != nil {
		t.Fatalf("error creating service: %v", err)
	}

	ctx := context.TODO()
	accountId := "accountid"

	account := &schemas.Account{}

	// expected error not found from DB ReadByID method
	mockDB.EXPECT().ReadByID(account, "id = ?", accountId).Return(gorm.ErrRecordNotFound)

	req := accountId

	res := s.GetAccountById(ctx, req)
	expectedResponse := response.BodyStruct{
		StatusCode: http.StatusNotFound,
		Message:    "account not found",
		Err:        gorm.ErrRecordNotFound,
	}

	assert.Equal(t, expectedResponse.StatusCode, res.StatusCode)
	assert.Equal(t, expectedResponse.Message, res.Message)
}

func TestCreateTransaction_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBHandlerInterface(ctrl)

	s, err := services.New(mockDB)
	if err != nil {
		t.Fatalf("error creating service: %v", err)
	}

	ctx := context.TODO()

	// set up expectations on the mock
	mockDB.EXPECT().Create(gomock.Any()).Return(nil)

	req := models.CreateTransactionRequest{
		AccountID:       "12344556",
		OperationTypeID: 1,
		Amount:          125.0,
	}

	res := s.CreateTransaction(ctx, req)
	expectedResponse := response.BodyStruct{
		StatusCode: http.StatusOK,
	}

	assert.Equal(t, expectedResponse.StatusCode, res.StatusCode)
}

func TestCreateTransaction_DBError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBHandlerInterface(ctrl)

	s, err := services.New(mockDB)
	if err != nil {
		t.Fatalf("error creating service: %v", err)
	}

	ctx := context.TODO()

	// mocking db error
	expectedErr := errors.New("failed to create transaction record in db")
	mockDB.EXPECT().Create(gomock.Any()).Return(expectedErr)

	req := models.CreateTransactionRequest{
		AccountID:       "12344556",
		OperationTypeID: 1,
		Amount:          125.0,
	}

	res := s.CreateTransaction(ctx, req)
	expectedResponse := response.BodyStruct{
		StatusCode: http.StatusInternalServerError,
		Message:    "failed to create transaction record in db",
	}

	assert.Equal(t, expectedResponse.StatusCode, res.StatusCode)
	assert.Equal(t, expectedResponse.Message, res.Message)
}
