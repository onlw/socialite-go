package provider

import (
	"bytes"
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/util"
	"net/url"
	"strings"
)

var baseURL = ""

type Wechat struct {
	Base
	config.Config
	ForcePopup bool
	Component  WechatComponent
}

type WechatComponent struct {
	ID    string
	Token string
}

func (p Wechat) GetAuthUrl() string {
	var buf bytes.Buffer
	path := "oauth2/authorize"
	if util.Contains("snsapi_login", p.Scopes) {
		path = "qrconnect"
	}
	base := "https://open.weixin.qq.com/connect/" + path

	// todo Add custom fields
	v := url.Values{
		"appid":         {p.ClientID},
		"redirect_uri":  {p.RedirectURL},
		"response_type": {"code"},
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

func (p Wechat) Redirect() {

}

func (p Wechat) User() contract.User {
	return model.User{}
}
