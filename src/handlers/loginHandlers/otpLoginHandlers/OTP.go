package otpLoginHandlers

import (
	"appauths/src/appTypes"
	"appauths/src/globalVars"
	"appauths/src/helpers"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ValidateOTP(c *fiber.Ctx) error {
	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	if session.Get("state").(string) != "login: 2FA with OTP" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var body struct {
		OTP int `json:"otp"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if session.Get("2faOTP").(int) != body.OTP {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("incorrect OTP")
	}

	user, err := helpers.QueryRowType[appTypes.User]("SELECT user_id, email, username FROM auth_user WHERE email = $1", session.Get("email").(string))
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := session.Destroy(); err != nil {
		panic(err)
	}

	// log the user in
	jwt := helpers.JwtSign(user, os.Getenv("AUTH_JWT_SECRET"), time.Now().Add(24*time.Hour))

	return c.JSON(fiber.Map{
		"msg":     "Login success!",
		"authJwt": jwt,
	})
}
