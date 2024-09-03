package loginHandlers

import (
	"appauths/backend/appTypes"
	"appauths/backend/helpers"
	"os"
	"time"

	"log"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(fiber.ErrUnprocessableEntity.Message)
	}

	user, err := helpers.QueryRowType[appTypes.User]("SELECT user_id, email, username FROM auth_user WHERE username = $1 AND password = $2", body.Username, body.Password)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if user == nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("incorrect username or password")
	}

	// create token -> (header.payload)
	jwt := helpers.JwtSign(user, os.Getenv("AUTH_JWT_SECRET"), time.Now().Add(24*time.Hour))

	return c.JSON(fiber.Map{
		"msg":     "Login success!",
		"authJwt": jwt,
	})
}
