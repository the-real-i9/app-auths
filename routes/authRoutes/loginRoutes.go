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

	router.Get("/totp/barcode_setupkey", totpLoginHandlers.BarcodeSetupKey)
	router.Post("/totp/verify", totpLoginHandlers.VerifyTOTP)
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
