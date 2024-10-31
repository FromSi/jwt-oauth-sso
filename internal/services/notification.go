package services

//go:generate mockgen -destination=../mocks/services/mock_query_notification_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services QueryNotificationService
type QueryNotificationService interface{}

//go:generate mockgen -destination=../mocks/services/mock_mutable_notification_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services MutableNotificationService
type MutableNotificationService interface {
	SendText(string)
}

//go:generate mockgen -destination=../mocks/services/mock_notification_service.go -package=services_mocks github.com/fromsi/jwt-oauth-sso/internal/services NotificationService
type NotificationService interface {
	QueryNotificationService
	MutableNotificationService
}
