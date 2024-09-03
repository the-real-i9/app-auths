package credSignupHandlers

import (
	"appauths/backend/globalVars"
	"appauths/backend/helpers"
	"time"

	"github.com/gofiber/fiber/v2"
)

func VerifyEmail(c *fiber.Ctx) error {
	session, err := globalVars.SignupSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	if session.Get("step").(string) != "verify email" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var body struct {
		InputVerfToken int `json:"verification_code"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("invalid payload")
	}

	email := session.Get("email").(string)
	verfToken := session.Get("verificationToken").(int)
	verfTokenExpires := time.Unix(session.Get("verificationTokenExpires").(int64), 0)

	if verfToken != body.InputVerfToken {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Incorrect verification code")
	}

	if time.Now().After(verfTokenExpires) {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("Verification code expired")
	}

	go helpers.SendMail(email, "Email verification success", "Your email has been verified!")

	session.Delete("verificationToken")
	session.Delete("verificationTokenExpires")
	session.Set("step", "register user")

	if save_err := session.Save(); save_err != nil {
		panic(save_err)
	}

	return c.SendString("Your email " + email + " has been verified!\n")
}
