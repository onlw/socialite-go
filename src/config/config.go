package config

type Config struct {
	ClientID string

	ClientSecret string

	Endpoint Endpoint

	RedirectURL string

	Scopes []string

	State string
}

type Endpoint struct {
	AuthURL string

	TokenURL string
}
