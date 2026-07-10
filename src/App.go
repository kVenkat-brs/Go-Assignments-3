package src

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kVenkat-brs/Go-Assignment-2/src/db"
	"github.com/kVenkat-brs/Go-Assignments-3/src/routes"
)

func SetupApp() *fiber.App {

	app := fiber.New()

	db.SetupDB()

	routes.Studentroutes(app)

	return app

}