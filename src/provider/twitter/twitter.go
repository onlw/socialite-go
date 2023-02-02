package github

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
)

const name = "twitter"

// https://developer.twitter.com/en/docs/authentication/oauth-2-0/authorization-code
var scopes = []string{
	"users.read",
	"tweet.read",
}

var endpoint = config.Endpoint{
	AuthURL: "https://twitter.com/i/oauth2/authorize",
}

type Twitter struct {
	provider.Base
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p Twitter) GetAuthURL() string {
	return p.GetAuthURLFromBase(endpoint.AuthURL)
}

func (p Twitter) Redirect() {

}

func (p Twitter) User() contract.User {
	return model.User{}
}
