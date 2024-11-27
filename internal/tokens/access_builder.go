package tokens

//go:generate mockgen -destination=../../mocks/tokens/mock_access_builder.go -package=tokens_mocks github.com/fromsi/jwt-oauth-sso/internal/tokens AccessTokenBuilder
type AccessTokenBuilder interface {
	New() AccessTokenBuilder
	NewFromJwtString(string) (AccessTokenBuilder, error)
	Build() (AccessToken, error)
	BuildToJwt() (*JwtAccessToken, error)
	SetIssuer(string) AccessTokenBuilder
	SetAudience(string) AccessTokenBuilder
	SetSubject(string) AccessTokenBuilder
	SetDeviceUUID(string) AccessTokenBuilder
	SetDeviceUserAgent(string) AccessTokenBuilder
	SetIssuedAt(int) AccessTokenBuilder
	SetExpirationTime(int) AccessTokenBuilder
}
