package app

import (
	"errors"
	"github.com/avtara/carehub/utils"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func (cfg *App) InitOauth() (err error) {
	var (
		clientID     = utils.GetEnv("oauth.google.client_id", "")
		clientSecret = utils.GetEnv("oauth.google.client_secret", "")
	)

	if clientID == "" || clientSecret == "" {
		return errors.New("client secret not found")
	}

	cfg.OauthConfig = &oauth2.Config{
		RedirectURL:  utils.GetEnv("oauth.google.redirect_url", "http://localhost:9000/callback"),
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	return nil
}
