package utils

import "github.com/gofiber/fiber/v3"

func Error(c fiber.Ctx, StatusCode int,message string) error{
	return c.Status(StatusCode).JSON(fiber.Map{
			"success": false,
			"message":message,
			})
}

func Success (c fiber.Ctx,StatusCode int,message string,data interface{}) error{
	return c.Status(StatusCode).JSON(fiber.Map{
		"success": true,
		"message":message,
		"User":data,
	})
}