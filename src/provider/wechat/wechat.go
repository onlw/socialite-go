package wechat

import (
	"bytes"
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
	"github.com/onlw/socialite-go/v2/util"
	"net/url"
	"strings"
)

const name = "wechat"

var scopes = []string{"snsapi_login"}

var endpoint = config.Endpoint{
	AuthURL: "https://open.weixin.qq.com/connect/",
}

type Wechat struct {
	provider.Base
	ForcePopup bool
	Component  Component
}

type Component struct {
	ID    string
	Token string
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p Wechat) GetAuthUrl() string {
	var buf bytes.Buffer
	path := "oauth2/authorize"
	if util.Contains("snsapi_login", p.Scopes) {
		path = "qrconnect"
	}
	base := endpoint.AuthURL + path

	// todo Add custom fields
	v := url.Values{
		"appid":         {p.ClientID},
		"redirect_uri":  {p.RedirectURL},
		"response_type": {"code"},
		"state":         {p.State},
	}
	v.Set("scope", strings.Join(p.Scopes, ","))
	if p.Component.ID != "" {
		v.Set("component_id", p.Component.ID)
	}
	buf.WriteString(base)
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	buf.WriteString("#wechat_redirect")

	return buf.String()
}

func (p Wechat) getScopes() []string {
	if len(p.Scopes) > 0 {
		return p.Scopes
	}

	return scopes
}

func (p Wechat) Redirect() {

}

func (p Wechat) User() contract.User {
	return model.User{}
}
