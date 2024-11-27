package tokens

//go:generate mockgen -destination=../../mocks/tokens/mock_query_access.go -package=tokens_mocks github.com/fromsi/jwt-oauth-sso/internal/tokens QueryAccessToken
type QueryAccessToken interface {
	ToString() (string, error)
	GetIssuer() string
	GetAudience() string
	GetSubject() string
	GetDeviceUUID() string
	GetDeviceUserAgent() string
	GetIssuedAt() int
	GetExpirationTime() int
}

//go:generate mockgen -destination=../../mocks/tokens/mock_mutable_access.go -package=tokens_mocks github.com/fromsi/jwt-oauth-sso/internal/tokens MutableAccessToken
type MutableAccessToken interface {
	SetIssuer(string)
	SetAudience(string)
	SetSubject(string)
	SetDeviceUUID(string)
	SetDeviceUserAgent(string)
	SetIssuedAt(int)
	SetExpirationTime(int)
}

//go:generate mockgen -destination=../../mocks/tokens/mock_access.go -package=tokens_mocks github.com/fromsi/jwt-oauth-sso/internal/tokens AccessToken
type AccessToken interface {
	QueryAccessToken
	MutableAccessToken
}
