package Gitee

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
)

const name = "gitee"

var scopes = []string{"user_info"}

var endpoint = config.Endpoint{
	AuthURL: "https://gitee.com/oauth/authorize",
}

type Gitee struct {
	provider.Base
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p Gitee) GetAuthURL() string {
	p.ScopeSeparator = ","

	return p.GetAuthURLFromBase(endpoint.AuthURL)
}

func (p Gitee) Redirect() {

}

func (p Gitee) User() contract.User {
	return model.User{}
}
