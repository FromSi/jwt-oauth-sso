package services

import "github.com/fromsi/jwt-oauth-sso/internal/repositories"

//go:generate mockgen -destination=../../mocks/services/mock_query_reset_token_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services QueryResetTokenService
type QueryResetTokenService interface {
	GenerateToken() string
}

//go:generate mockgen -destination=../../mocks/services/mock_mutable_reset_token_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services MutableResetTokenService
type MutableResetTokenService interface {
	SendNewResetTokenByUser(repositories.User) error
}

//go:generate mockgen -destination=../../mocks/services/mock_reset_token_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services ResetTokenService
type ResetTokenService interface {
	QueryResetTokenService
	MutableResetTokenService
}
