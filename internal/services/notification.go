package services

import "github.com/fromsi/jwt-oauth-sso/internal/repositories"

//go:generate mockgen -destination=../../mocks/services/mock_query_notification.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services QueryNotificationService
type QueryNotificationService interface{}

//go:generate mockgen -destination=../../mocks/services/mock_mutable_notification.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services MutableNotificationService
type MutableNotificationService interface {
	SendTextByUser(repositories.User, string) error
}

//go:generate mockgen -destination=../../mocks/services/mock_notification.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services NotificationService
type NotificationService interface {
	QueryNotificationService
	MutableNotificationService
}
