package provider

import (
	"bytes"
	"github.com/onlw/socialite-go/v2/config"
	"net/http"
	"net/url"
	"strings"
)

type Base struct {
	config.Config
}

func (p Base) RedirectTo() {
	http.RedirectHandler(p.GetAuthURL(), http.StatusFound)
}

func (p Base) GetAuthURL() string {
	return ""
}

func RegisterProvider(provider string, defaultScopes []string) {
	// todo add RegisterProvider
}

func (p Base) GetAuthURLFromBase(authURL string) string {
	var buf bytes.Buffer
	buf.WriteString(authURL)
	v := p.GetAuthCodeFields()
	//for _, opt := range opts {
	//	opt.setValue(v)
	//}
	if strings.Contains(p.Endpoint.AuthURL, "?") {
		buf.WriteByte('&')
	} else {
		buf.WriteByte('?')
	}
	buf.WriteString(v.Encode())

	return buf.String()
}

func (p Base) GetAuthCodeFields() url.Values {
	v := url.Values{
		"client_id":     {p.ClientID},
		"response_type": {"code"},
	}
	if p.RedirectURL != "" {
		v.Set("redirect_uri", p.RedirectURL)
	}
	if len(p.Scopes) > 0 {
		v.Set("scope", strings.Join(p.Scopes, " "))
	}
	if p.State != "" {
		v.Set("state", p.State)
	}

	return v
}
