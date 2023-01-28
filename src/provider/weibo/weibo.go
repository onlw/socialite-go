package WeiBo

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
)

const name = "weibo"

var scopes = []string{"email"}

var endpoint = config.Endpoint{
	AuthURL: "https://api.weibo.com/oauth2/authorize",
}

type WeiBo struct {
	provider.Base
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p WeiBo) GetAuthURL() string {
	p.ScopeSeparator = ","

	return p.GetAuthURLFromBase(endpoint.AuthURL)
}

func (p WeiBo) Redirect() {

}

func (p WeiBo) User() contract.User {
	return model.User{}
}
