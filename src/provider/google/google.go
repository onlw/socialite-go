package github

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
)

const name = "google"

var scopes = []string{
	"https://www.googleapis.com/auth/userinfo.email",
	"https://www.googleapis.com/auth/userinfo.profile",
}

var endpoint = config.Endpoint{
	AuthURL: "https://accounts.google.com/o/oauth2/v2/auth",
}

type Github struct {
	provider.Base
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p Github) GetAuthURL() string {
	return p.GetAuthURLFromBase(endpoint.AuthURL)
}

func (p Github) Redirect() {

}

func (p Github) User() contract.User {
	return model.User{}
}
