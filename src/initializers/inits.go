package initializers

import (
	"appauths/src/globalVars"
	"context"
	"os"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/oauth2"
)

func InitOauth2Config() {
	globalVars.Oauth2Config = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "openid"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://oauth2.googleapis.com/token",
		},
		RedirectURL: "http://127.0.0.1:5000/api/auth/oauth/google/callback",
	}
}

func InitDBPool() error {
	pool, err := pgxpool.New(context.Background(), os.Getenv("PGDATABASE_URL"))
	if err != nil {
		return err
	}
	globalVars.DBPool = pool

	return nil
}

func InitSessionStores() {
	authStorage := postgres.New(postgres.Config{ConnectionURI: os.Getenv("PGDATABASE_URL"), Table: "ongoing_auth"})
	appStorage := postgres.New(postgres.Config{ConnectionURI: os.Getenv("PGDATABASE_URL"), Table: "ongoing_process"})

	globalVars.AuthSessionStore = session.New(session.Config{
		Storage:    authStorage,
		CookiePath: "/api/auth",
	})

	globalVars.AppSessionStore = session.New(session.Config{
		Storage:    appStorage,
		CookiePath: "/api/app",
	})
}
