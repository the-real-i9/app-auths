package globalVars

import (
	"errors"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/oauth2"
)

var AuthSessionStore *session.Store

var AppSessionStore *session.Store

var GoogleOauth2Config *oauth2.Config

var GithubOauth2Config *oauth2.Config

var ErrInternalServerError = errors.New("internal server error: check logger")

var DBPool *pgxpool.Pool
