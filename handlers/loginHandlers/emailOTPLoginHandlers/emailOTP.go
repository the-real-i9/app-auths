package emailOTPLoginHandlers

import (
	"appauths/helpers"
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

func SendOTP(c *fiber.Ctx) error {

	var body struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(fiber.ErrUnprocessableEntity.Message)
	}

	otp := rand.Intn(899999) + 100000

	go helpers.SendMail(body.Email, "2FA Login OTP", fmt.Sprintf("Your Login OTP (One-Time Password) is %d", otp))

	return c.Status(200).SendString("OTP has been sent to " + body.Email)
}

func VerifyOTP(c *fiber.Ctx) error {

	return nil
}
