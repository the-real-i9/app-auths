package otpLoginHandlers

import (
	"appauths/src/appTypes"
	"appauths/src/globalVars"
	"appauths/src/helpers"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

// this implementation assumes an existing login session coming from credential login
func SendOTP(c *fiber.Ctx) error {
	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	if session.Get("state").(string) != "login: 2FA with OTP" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	email := session.Get("email").(string)

	otp := rand.Intn(899999) + 100000

	go helpers.SendMail(email, "2FA Login OTP", fmt.Sprintf("Your Login OTP (One-Time Password) is %d", otp))

	session.Set("loginOTP", otp)
	session.Set("state", "otp_login: verify otp")

	if save_err := session.Save(); save_err != nil {
		panic(save_err)
	}

	return c.Status(200).SendString("Login OTP has been sent to " + email)
}

func VerifyOTP(c *fiber.Ctx) error {
	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	if session.Get("state").(string) != "otp_login: verify otp" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var body struct {
		OTP int `json:"otp"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	if session.Get("loginOTP").(int) != body.OTP {
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
