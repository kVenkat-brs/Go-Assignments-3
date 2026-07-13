package utils

import "github.com/gofiber/fiber/v3"

func Error(c fiber.Ctx, Status int,message string) error{
	return c.Status(Status).JSON(fiber.Map{
			"success": false,
			"message":message,
			})
}

func Success (c fiber.Ctx,Status int,message string,data interface{}) error{
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message":message,
		"User":data,
	})
}