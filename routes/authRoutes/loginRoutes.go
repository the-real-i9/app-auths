package authRoutes

import (
	"appauths/handlers/loginHandlers"
	"appauths/handlers/loginHandlers/emailOTPLoginHandlers"
	"appauths/handlers/loginHandlers/totpLoginHandlers"

	"github.com/gofiber/fiber/v2"
)

// Multi-factor authentication
func MFALogin(router fiber.Router) {
	router.Post("/otp/email/send_otp", emailOTPLoginHandlers.SendOTP)
	router.Post("/otp/email/verify", emailOTPLoginHandlers.VerifyOTP)

	// create session and store secret in it for the next endpoint
	router.Get("/totp/setup/barcode_setupkey", totpLoginHandlers.BarcodeSetupKey)
	// retrieve secret from session,
	// if it works, store it permanently,
	// destroy session
	router.Post("/totp/setup/validate_passcode", totpLoginHandlers.ValidatePasscodeSetup)

	router.Post("/totp/validate_passcode", totpLoginHandlers.ValidatePasscode)
}

// SSO login
func SSO(router fiber.Router) {

}

func Login(router fiber.Router) {
	// a cookie restricted to this path is used to maintain a session through the login process
	router.Post("/cred", loginHandlers.CredLogin)

	router.Route("/mfa", MFALogin)

	router.Route("/sso", SSO)
}
