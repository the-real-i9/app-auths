package totpLoginHandlers

import (
	"appauths/src/appTypes"
	"appauths/src/globalVars"
	"appauths/src/helpers"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
)

// this implementation assumes an existing login session coming from credential login
func ValidatePasscode(c *fiber.Ctx) error {
	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	if session.Get("state").(string) != "login: 2FA with TOTP" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var body struct {
		Passcode string `json:"passcode"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	email := session.Get("email").(string)

	setupKey, dberr := helpers.QueryRowField[string]("SELECT totp_setup_key FROM auth_user WHERE email = $1", email)
	if dberr != nil {
		log.Println(dberr)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if valid := totp.Validate(body.Passcode, *setupKey); !valid {
		return c.Status(fiber.StatusUnauthorized).SendString("passcode incorrect")
	}

	user, err := helpers.QueryRowType[appTypes.User]("SELECT user_id, email, username FROM auth_user WHERE email = $1", email)
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
