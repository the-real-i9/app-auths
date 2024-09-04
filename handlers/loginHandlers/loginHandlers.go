package loginHandlers

import (
	"appauths/appTypes"
	"appauths/helpers"
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

	userData, err := helpers.QueryRowType[map[string]any]("SELECT user_id, email, username, password FROM auth_user WHERE username = $1", body.Username)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if userData == nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("incorrect username or password")
	}

	hashedPwd := (*userData)["password"].(string)

	if cmp_err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(body.Password)); cmp_err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("incorrect username or password")
	}

	var user appTypes.User

	helpers.MapToStruct(*userData, &user)

	jwt := helpers.JwtSign(user, os.Getenv("AUTH_JWT_SECRET"), time.Now().Add(24*time.Hour))

	return c.JSON(fiber.Map{
		"msg":     "Login success!",
		"authJwt": jwt,
	})
}
