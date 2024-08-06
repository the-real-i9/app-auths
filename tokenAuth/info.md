# Token Authentication Strageties

JWTs are used to repeatedly authenticate users into your API.

OTPs (One Time Passwords) are tokens generally used to confirm that a user attempting to perform an operation is who he claims he is, before allowing a restricted operation, thus ensuring secure access. Common use cases are:

- email or phone number verification (before or after user signup): you don't want users to sign-up with someone else's email or one that doesn't exist, either you prevent the sign-up or you limit user priviledges
- re-logins: additional layer of security (2FA), or perhaps, after many failed login attempts
- account recovery: since you forgot your account key/secret, here's another way we can confirm its really you
- password reset: preventing attackers from taking control of an account and rendering it inaccessible to the original user
- payment confirmation: even if an attacker has stolen the user's password, you don't want your organization to be sued for a critical operation like this one. It's important to ask, "Are you the one making this payment, or someone else?"
