package totpLoginHandlers

import (
	"appauths/globalVars"
	"appauths/helpers"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func generateBarcodeAndSetupKey(accName string) (barcodeImageURL, setupKey string) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "i9appauths",
		AccountName: accName,
		Algorithm:   otp.AlgorithmSHA256,
		SecretSize:  8,
	})
	if err != nil {
		panic(err)
	}

	qrImage, err := key.Image(200, 200)
	if err != nil {
		panic(err)
	}

	var qrImageData bytes.Buffer

	if err := png.Encode(&qrImageData, qrImage); err != nil {
		panic(err)
	}

	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(qrImageData.Bytes())), key.Secret()
}

func BarcodeSetupKey(c *fiber.Ctx) error {

	var body struct {
		Email string `json:"email"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	barcodeImageURL, setupKey := generateBarcodeAndSetupKey(body.Email)

	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	session.Set("step", "totp auth setup: validate passcode")
	session.Set("accName", body.Email)
	session.Set("setupKey", setupKey)
	session.SetExpiry(30 * time.Minute)

	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.JSON(fiber.Map{
		"barcodeImageURL": barcodeImageURL,
		"setupKey":        setupKey,
	})
}

func ValidatePasscodeSetup(c *fiber.Ctx) error {
	session, err := globalVars.AuthSessionStore.Get(c)
	if err != nil {
		panic(err)
	}

	if session.Get("step").(string) != "totp auth setup: validate passcode" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	var body struct {
		Passcode string `json:"passcode"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).SendString(err.Error())
	}

	setupKey := session.Get("setupKey").(string)

	if valid := totp.Validate(body.Passcode, setupKey); !valid {
		return c.Status(fiber.StatusUnprocessableEntity).SendString("totp setup fail: passcode or setup key incorrect")
	}

	email := session.Get("accName").(string)

	_, dberr := helpers.QueryRowField[bool]("UPDATE auth_user SET totp_setup_key = $1 WHERE email = $2 RETURNING true", setupKey, email)
	if dberr != nil {
		log.Println(dberr)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).SendString("TOTP setup success")
}

func ValidatePasscode(c *fiber.Ctx) error {

	return nil
}
