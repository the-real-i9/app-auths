package emailOTPLoginHandlers

import (
	"appauths/globalVars"
	"appauths/helpers"
	"fmt"
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SendOTP(c *fiber.Ctx) error {

	var body struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	otp := rand.Intn(899999) + 100000

	go helpers.SendMail(body.Email, "2FA Login OTP", fmt.Sprintf("Your Login OTP (One-Time Password) is %d", otp))

	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	session.Set("email", body.Email)
	session.Set("loginOTP", otp)
	session.Set("step", "email_otp_login: verify otp")
	session.SetExpiry(5 * time.Minute)

	if save_err := session.Save(); save_err != nil {
		panic(save_err)
	}

	return c.Status(200).SendString("Login OTP has been sent to " + body.Email)
}

func VerifyOTP(c *fiber.Ctx) error {

	var body struct {
		OTP int `json:"otp"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	if session.Get("step").(string) != "email_otp_login: verify otp" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if session.Get("loginOTP").(int) != body.OTP {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("incorrect OTP")
	}

	if err := session.Destroy(); err != nil {
		panic(err)
	}

	return c.SendStatus(200)
}
