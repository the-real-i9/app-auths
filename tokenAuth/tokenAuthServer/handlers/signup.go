package handlers

import "github.com/gofiber/fiber/v2"

func submitEmailHandler(ctx *fiber.Ctx) error {

	return nil
}

func verifyEmailHandler(ctx *fiber.Ctx) error {

	return nil
}

func registerUserHandler(ctx *fiber.Ctx) error {

	return nil
}

func Signup(router fiber.Router) {
	router.Post("/submit_email", submitEmailHandler)
	router.Post("/verify_email", verifyEmailHandler)
	router.Post("/register_user", registerUserHandler)
}
