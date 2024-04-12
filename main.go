package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/chandan782/Pismo/api/controllers"
	"github.com/chandan782/Pismo/configs"
	"github.com/chandan782/Pismo/db"
	_ "github.com/chandan782/Pismo/docs"
	"github.com/chandan782/Pismo/routes"
	"github.com/gofiber/fiber/v2"
)

// @title Pismo API
// @description This is the Pismo API for managing user accounts and transactions.
// @version 1.0
// @host localhost:3000
// @BasePath /api/v1
func main() {
	// create a new Fiber instance
	app := fiber.New()

	// load server configs
	serverCfg := configs.GetServerConfig()

	// initialize the database
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.CloseDB()

	// initialize the controller instance
	c, err := controllers.New(&db.DBHandler{DB: db.DB})
	if err != nil {
		log.Fatalf("Error initializing controllers: %v", err)
	}

	// define routes
	routes.SetupRoutes(app, c)

	// Register route for Swagger UI (assuming docs directory is next to main.go)
	routes.SetupSwagger(app)
	// app.Get("/swagger/*", swagger.HandlerDefault)

	// create a channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// start the server
	go func() {
		err := app.Listen(":" + serverCfg.Port)
		if err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	log.Printf("server is running on port %s", serverCfg.Port)

	// Wait for OS signal to gracefully shutdown the server
	<-quit
	log.Println("shutting down server...")
}
