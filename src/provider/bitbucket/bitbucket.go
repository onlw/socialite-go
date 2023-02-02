package github

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
)

const name = "Bitbucket"

// https://developer.Bitbucket.com/en/docs/authentication/oauth-2-0/authorization-code
var scopes = []string{
	"users.read",
	"tweet.read",
}

var endpoint = config.Endpoint{
	AuthURL: "https://Bitbucket.com/i/oauth2/authorize",
}

type Bitbucket struct {
	provider.Base
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p Bitbucket) GetAuthURL() string {
	return p.GetAuthURLFromBase(endpoint.AuthURL)
}

func (p Bitbucket) Redirect() {

}

func (p Bitbucket) User() contract.User {
	return model.User{}
}
