package integration_test

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/chandan782/Pismo/api/controllers"
	"github.com/chandan782/Pismo/api/models"
	"github.com/chandan782/Pismo/configs"
	"github.com/chandan782/Pismo/db"
	"github.com/chandan782/Pismo/response"
	"github.com/chandan782/Pismo/routes"
	"github.com/pkg/errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	dot "github.com/joho/godotenv"
)

var (
	app        *fiber.App
	client     http.Client
	documentNo = "1234567890"
)

func TestMain(m *testing.M) {
	// Initialize database connection
	err := dot.Load(".env.sample")
	if err != nil {
		panic(err)
	}

	err = db.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.CloseDB()

	// Load server configs
	serverCfg := configs.GetServerConfig()

	// Create a new Fiber instance
	app = fiber.New()

	// Initialize controllers
	dbHandler := db.DBHandler{DB: db.DB}
	c, err := controllers.New(&dbHandler)
	if err != nil {
		panic(err)
	}

	// Define routes
	routes.SetupRoutes(app, c)

	// Wrap the app with recover middleware (optional for error handling)
	app.Use(recover.New())

	client = http.Client{
		Timeout: 30 * time.Second,
	}

	// Start the server in a separate goroutine for testing purposes
	go func() {
		err := app.Listen(":" + serverCfg.Port)
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Run tests
	exitVal := m.Run()

	os.Exit(exitVal)
}

func createAccountApiCalled() (models.CreateAccountResponse, error) {
	reqBody := fmt.Sprintf(`{
		"document_number": "%s"
	}`, documentNo)

	// Perform the POST request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%s/api/v1/accounts/", configs.GetServerConfig().Port), strings.NewReader(reqBody))
	if err != nil {
		return models.CreateAccountResponse{}, errors.Errorf("failed to create new request: %v", req)
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return models.CreateAccountResponse{}, errors.Errorf("error sending request: %v", err)
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		return models.CreateAccountResponse{}, errors.Errorf("error reading response body: %v", err)
	}

	var resBody response.BodyStruct
	err = json.Unmarshal(resByte, &resBody)
	if err != nil {
		return models.CreateAccountResponse{}, errors.Errorf("error unmarshalling response body: %v", err)
	}

	dataByte, err := json.Marshal(resBody.Data)
	if err != nil {
		return models.CreateAccountResponse{}, errors.Errorf("failed to marshal response body data: %v", err)
	}

	var createAccountResponse models.CreateAccountResponse
	err = json.Unmarshal(dataByte, &createAccountResponse)
	if err != nil {
		return models.CreateAccountResponse{}, errors.Errorf("error unmarshalling response body data byte: %v", err)
	}

	return createAccountResponse, nil
}

func TestCreateAccount(t *testing.T) {
	createAccountRes, err := createAccountApiCalled()
	if err != nil {
		t.Errorf("%v", err)
	}

	// check if ID is generated and not empty
	if createAccountRes.ID == "" {
		t.Errorf("expected non-empty ID in response, got empty")
	}
}

func TestGetAccountById(t *testing.T) {
	// call create account first
	createAccountRes, err := createAccountApiCalled()
	if err != nil {
		t.Errorf("%v", err)
	}

	// perform GET request
	reqURL := fmt.Sprintf("http://localhost:%s/api/v1/accounts/%s", configs.GetServerConfig().Port, createAccountRes.ID)
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("error sending request: %v", err)
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	var resBody response.BodyStruct
	err = json.Unmarshal(resByte, &resBody)
	if err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}

	dataByte, err := json.Marshal(resBody.Data)
	if err != nil {
		t.Errorf("failed to marshal response body data: %v", err)
	}

	var getAcctByIdRes models.GetAccountByIdResponse
	err = json.Unmarshal(dataByte, &getAcctByIdRes)
	if err != nil {
		t.Errorf("error unmarshalling response body data byte: %v", err)
	}

	// Check if retrieved account details match created account
	if getAcctByIdRes.ID != createAccountRes.ID || getAcctByIdRes.DocumentNumber != documentNo {
		t.Errorf("Retrieved account details don't match created account")
	}
}

func TestCreateTransaction(t *testing.T) {
	// call create account first
	createAccountRes, err := createAccountApiCalled()
	if err != nil {
		t.Errorf("%v", err)
	}

	reqBody := fmt.Sprintf(`{
		"account_id": "%s",
		"operation_type_id": 2,
		"amount": 123.45
		}`, createAccountRes.ID)

	// perform the POST request
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://localhost:%s/api/v1/transactions/", configs.GetServerConfig().Port), strings.NewReader(reqBody))
	if err != nil {
		t.Errorf("Error creating request: %v", err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("error sending request: %v", err)
	}

	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error reading response body: %v", err)
	}

	var resBody response.BodyStruct
	err = json.Unmarshal(resByte, &resBody)
	if err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}

	dataByte, err := json.Marshal(resBody.Data)
	if err != nil {
		t.Errorf("failed to marshal response body data: %v", err)
	}

	var createTxnRes models.CreateTransactionResponse
	err = json.Unmarshal(dataByte, &createTxnRes)
	if err != nil {
		t.Errorf("error unmarshalling response body data byte: %v", err)
	}

	// Check if ID is generated and not empty
	if createTxnRes.ID == "" {
		t.Errorf("Expected non-empty ID in response, got empty")
	}
}
