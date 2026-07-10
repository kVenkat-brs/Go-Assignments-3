package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/kVenkat-brs/Go-Assignment-2/src/Controllers"
	"github.com/kVenkat-brs/Go-Assignments-3/src/middlewares"
	ctrA3 "github.com/kVenkat-brs/Go-Assignments-3/src/controllers"
)

func Studentroutes(app *fiber.App) {
		app.Use(logger.New())
		app.Use(cors.New(cors.Config{AllowOrigins: []string{"http://localhost:5173"}}))
		app.Use(middlewares.RequestTimingMiddleware)
	grp:=app.Group("/students")
	grp.Get("/",middlewares.ApiKeyMiddleware,middlewares.AuthorizationMiddleware,controllers.GetAllstudents)
	grp.Get("/:id",middlewares.ApiKeyMiddleware,middlewares.AuthorizationMiddleware,controllers.GetAllstudents)
	grp.Post("/",middlewares.ApiKeyMiddleware,middlewares.AuthorizationMiddleware,controllers.CreateStudent)
	grp.Put("/:id",middlewares.ApiKeyMiddleware,middlewares.AuthorizationMiddleware,controllers.UpdateStudents)
	grp.Delete("/:id",middlewares.ApiKeyMiddleware,middlewares.AuthorizationMiddleware,controllers.DeleteStudent)
	grp.Post("/users",ctrA3.CreateUser)
	grp.Post("/login",ctrA3.Login)

}