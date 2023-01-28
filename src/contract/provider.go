package contract

type Provider interface {
	RedirectTo()

	GetAuthURL() string

	User() User
}
