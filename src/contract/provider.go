package contract

type Provider interface {
	RedirectTo()

	GetAuthUrl() string

	User() User
}
