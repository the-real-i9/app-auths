package loginHandlers

import (
	"appauths/appTypes"
	"appauths/globalVars"
	"appauths/helpers"
	"fmt"
	"math/rand"
	"os"
	"time"

	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Login with credentials
func CredLogin(c *fiber.Ctx) error {

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(fiber.ErrUnprocessableEntity.Message)
	}

	userData, err := helpers.QueryRowType[map[string]any]("SELECT user_id, email, username, password, mfa_enabled, mfa_type FROM auth_user WHERE username = $1", body.Username)

	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if userData == nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("incorrect username or password")
	}

	hashedPwd := (*userData)["password"].(string)

	if cmp_err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(body.Password)); cmp_err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("incorrect username or password")
	}

	var user appTypes.User

	helpers.MapToStruct(*userData, &user)

	mfaEnabled := (*userData)["mfa_enabled"].(bool)

	if !mfaEnabled {
		jwt := helpers.JwtSign(user, os.Getenv("AUTH_JWT_SECRET"), time.Now().Add(24*time.Hour))

		return c.JSON(fiber.Map{
			"2faEnabled": false,
			"msg":        "Login success!",
			"authJwt":    jwt,
		})
	}

	// 2FA is enabled for the user
	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	// set "state" key to either "login: 2FA with TOTP" or "login: 2FA with Email OTP", whichever you choose
	// session.Set("state", "login: 2FA with Email OTP")
	mfaType := (*userData)["mfa_type"].(string)

	switch mfaType {
	case "totp":
		session.Set("state", "login: 2FA with TOTP")
	case "otp":
		otp := rand.Intn(899999) + 100000

		// send OTP
		go helpers.SendMail(user.Email, "2FA Login OTP", fmt.Sprintf("Your Login OTP (One-Time Password) is %d", otp))

		session.Set("state", "login: 2FA with OTP")
		session.Set("2faOTP", otp)
	default:
	}

	session.Set("email", user.Email)
	session.SetExpiry(30 * time.Minute)

	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"2faEnabled": true,
		"2faType":    mfaType,
		"msg":        "Proceed to the target 2FA type endpoint",
	})
}
