package wework

import (
	"bytes"
	"github.com/onlw/socialite-go/v2/config"
	"github.com/onlw/socialite-go/v2/contract"
	"github.com/onlw/socialite-go/v2/model"
	"github.com/onlw/socialite-go/v2/provider"
	"net/url"
	"strings"
)

const name = "wework"

var scopes = []string{"snsapi_privateinfo"}

var endpoint = config.Endpoint{
	AuthURL: "https://open.weixin.qq.com/connect/oauth2/authorize",
}

var QRCodeAuthURL = "https://open.work.weixin.qq.com/wwopen/sso/qrConnect"

type WeWork struct {
	provider.Base
	UseQRCode bool
}

func init() {
	provider.RegisterProvider(name, scopes)
}

func (p WeWork) GetAuthURL() string {
	authURL := endpoint.AuthURL
	v := url.Values{
		"appid":         {p.ClientID},
		"agentid":       {p.ClientID},
		"redirect_uri":  {p.RedirectURL},
		"response_type": {"code"},
		"state":         {p.State},
	}
	v.Set("scope", strings.Join(p.Scopes, ","))

	if p.UseQRCode {
		authURL = QRCodeAuthURL
		v.Del("scope")
	}
	var buf bytes.Buffer
	buf.WriteString(authURL)
	buf.WriteByte('?')
	buf.WriteString(v.Encode())
	buf.WriteString("#wechat_redirect")

	return buf.String()
}

func (p WeWork) Redirect() {

}

func (p WeWork) User() contract.User {
	return model.User{}
}
