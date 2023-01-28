package contract

type User interface {
	GetId() string

	GetNickname() string

	GetName() string

	GetEmail() string

	GetAvatar() string

	GetRaw() string

	GetAccessToken()

	GetRefreshToken()

	GetExpiresIn()

	GetProvider()

	//setRefreshToken(?string $refreshToken)
	//
	//setExpiresIn(int $expiresIn)
	//
	//setTokenResponse(array $response)
	//
	//getTokenResponse()
	//
	//setProvider(ProviderInterface $provider)
	//
	//getRaw()
	//
	//setRaw(array $user)
	//
	//setAccessToken(string $token)
}
