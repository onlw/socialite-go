package qq

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
)

const name = "qq"

var scopes = []string{"get_user_info"}

var endpoint = config.Endpoint{
	AuthURL: "https://graph.qq.com/oauth2.0/authorize",
}

type QQ struct {
	provider.Base
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p QQ) GetAuthURL() string {
	return p.GetAuthURLFromBase(endpoint.AuthURL)
}

func (p QQ) Redirect() {

}

func (p QQ) User() contract.User {
	return model.User{}
}
