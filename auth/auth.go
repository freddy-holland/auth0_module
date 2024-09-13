package auth

import (
	"context"
	"errors"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

var SessionStore *sessions.CookieStore

func NewStore(secret string) {
	SessionStore = sessions.NewCookieStore([]byte(secret))
}

type AuthConfig struct {
	Auth0Domain       string
	Auth0ClientID     string
	Auth0ClientSecret string
	Auth0CallbackURL  string
}

type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

var Auth *Authenticator

func NewAuth(cfg *AuthConfig) error {
	provider, err := oidc.NewProvider(
		context.Background(),
		"https://"+cfg.Auth0Domain+"/",
	)

	if err != nil {
		return err
	}

	conf := oauth2.Config{
		ClientID:     cfg.Auth0ClientID,
		ClientSecret: cfg.Auth0ClientSecret,
		RedirectURL:  cfg.Auth0CallbackURL,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile"},
	}

	Auth = &Authenticator{
		Provider: provider,
		Config:   conf,
	}

	return nil
}

func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}
