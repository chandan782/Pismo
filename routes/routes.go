package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/chandan782/Pismo/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, c controllers.Controllers) {
	// account routes
	app.Post("/api/v1/accounts/", c.CreateAccount)
	app.Get("/api/v1/accounts/:id", c.GetAccountById)

	// transactions routes
	app.Post("/api/v1/transactions/", c.CreateTransaction)
}

func SetupSwagger(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
