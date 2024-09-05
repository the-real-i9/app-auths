package routes

import (
	"appauths/appTypes"
	"appauths/handlers/totpHandlers"
	"appauths/helpers"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func App(router fiber.Router) {
	router.Use(func(c *fiber.Ctx) error {
		authToken := strings.Fields(c.GetReqHeaders()["Authorization"][0])[1]

		user, err := helpers.JwtVerify[appTypes.User](authToken, os.Getenv("AUTH_JWT_SECRET"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		c.Locals("user", user)

		return c.Next()
	})

	// access a restricted resource : jwt auth
	router.Get("/restricted", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*appTypes.User)

		return c.JSON(user)
	})
	router.Get("/totp/setup/barcode_setupkey", totpHandlers.BarcodeSetupKey)
	router.Post("/totp/setup/validate_passcode", totpHandlers.ValidateSetupPasscode)
}
