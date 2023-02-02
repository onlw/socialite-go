package github

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
)

const name = "linkedin"

var scopes = []string{
	"r_liteprofile",
	"r_emailaddress",
}

var endpoint = config.Endpoint{
	AuthURL: "https://www.linkedin.com/oauth/v2/authorization",
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
