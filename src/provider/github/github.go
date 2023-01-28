package github

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
)

const name = "github"

var scopes = []string{"read:user"}

var endpoint = config.Endpoint{
	AuthURL: "https://github.com/login/oauth/authorize",
}

type Github struct {
	provider.Base
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p Github) GetAuthUrl() string {
	return p.GetAuthUrlFromBase(endpoint.AuthURL)
}

func (p Github) Redirect() {

}

func (p Github) User() contract.User {
	return model.User{}
}
