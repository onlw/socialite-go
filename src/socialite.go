package socialite

import "github.com/onlw/socialite-go/v2/contract"

type Socialite struct {
	providers map[string]contract.Provider
}

func New() Socialite {
	return Socialite{}
}

func (s Socialite) Provider(p string) contract.Provider {
	return s.providers[p]
}
