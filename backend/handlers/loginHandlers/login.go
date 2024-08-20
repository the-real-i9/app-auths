package loginHandlers

import (
	"i9codesauths/backend/helpers"
	"os"
	"time"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *fiber.Ctx) error {

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("invalid payload")
	}

	type User struct {
		Id       int    `db:"user_id" json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	user, err := helpers.QueryRowType[User]("SELECT user_id, email, username FROM i9ca_user WHERE username = $1 AND password = $2", body.Username, body.Password)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if user == nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("incorrect username or password")
	}

	// create token -> (header.payload)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data":  *user,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	// sign token with secret -> (header.payload.signature)
	jwt, err := token.SignedString([]byte(os.Getenv("AUTH_JWT_SECRET")))
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"msg": "Login success!",
		"jwt": jwt,
	})
}
