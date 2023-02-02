package facebook

import (
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
	"net/url"
	"strings"
)

const name = "facebook"

var scopes = []string{"email"}

var endpoint = config.Endpoint{
	AuthURL: "https://www.facebook.com/v15.0/dialog/oauth",
}

type Facebook struct {
	provider.Base
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p Facebook) GetAuthURL() string {
	return p.GetAuthURLFromBase(endpoint.AuthURL)
}

func (p Facebook) GetAuthCodeFields() url.Values {
	v := url.Values{
		"client_id":     {p.ClientID},
		"response_type": {"code"},
	}
	if p.RedirectURL != "" {
		v.Set("redirect_uri", p.RedirectURL)
	}
	if len(p.Scopes) > 0 {
		v.Set("scope", strings.Join(p.Scopes, p.ScopeSeparator))
	}
	if p.State != "" {
		v.Set("state", p.State)
	}

	return v
}

func (p Facebook) Redirect() {

}

func (p Facebook) User() contract.User {
	return model.User{}
}
