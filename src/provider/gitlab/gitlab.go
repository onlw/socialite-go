package github

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
)

const name = "gitlab"

var scopes = []string{"read_user"}

var endpoint = config.Endpoint{
	AuthURL: "https://gitlab.com/oauth/authorize",
}

type Github struct {
	provider.Base
	Host string
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p Github) GetAuthURL() string {
	if p.Host == "" {
		return p.GetAuthURLFromBase(endpoint.AuthURL)
	}

	return p.GetAuthURLFromBase(p.Host + "/oauth/authorize")
}

func (p Github) Redirect() {

}

func (p Github) User() contract.User {
	return model.User{}
}
