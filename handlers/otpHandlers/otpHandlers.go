package otpHandlers

import (
	"appauths/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func EnableOTP2FA(c *fiber.Ctx) error {
	var body struct {
		Username string `json:"username"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	_, err := helpers.QueryRowField[bool]("UPDATE auth_user SET mfa_enabled = $1, mfa_type $2 WHERE username = $3 RETURNING true", true, "otp", body.Username)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendString("OTP 2FA enabled")
}
