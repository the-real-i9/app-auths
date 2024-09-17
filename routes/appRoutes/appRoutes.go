package appRoutes

import (
	"appauths/appTypes"
	"appauths/handlers/otpHandlers"
	"appauths/handlers/totpHandlers"
	"appauths/helpers"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func App(router fiber.Router) {
	router.Use(func(c *fiber.Ctx) error {
		authHeader := c.GetReqHeaders()["Authorization"]

		if len(authHeader) == 0 {
			return c.Status(fiber.StatusUnauthorized).SendString("authorization required")
		}

		authToken := strings.Fields(authHeader[0])[1]

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

	router.Get("/totp_2fa/setup/barcode_setupkey", totpHandlers.BarcodeSetupKey)
	router.Post("/totp_2fa/setup/validate_passcode", totpHandlers.ValidateSetupPasscode)

	router.Put("/otp_2fa/enable", otpHandlers.EnableOTP2FA)
}
