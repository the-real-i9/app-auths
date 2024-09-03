package signupHandlers

import (
	"appauths/backend/globalVars"
	"appauths/backend/helpers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	session, err := globalVars.SignupSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	if session.Get("step").(string) != "register user" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("invalid payload")
	}

	// check if user with username already exists
	userExists, err := helpers.QueryRowField[bool]("SELECT EXISTS(SELECT 1 FROM auth_user WHERE username = $1)", body.Username)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if *userExists {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("username already taken")
	}

	email := session.Get("email").(string)

	_, dbin_err := helpers.QueryRowField[bool]("INSERT INTO auth_user (email, username, password) VALUES ($1, $2, $3) RETURNING true", email, body.Username, body.Password)
	if dbin_err != nil {
		log.Println(dbin_err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if sd_err := session.Destroy(); sd_err != nil {
		panic(sd_err)
	}

	return c.SendString("Registration success. Proceed to login.\n")
}
