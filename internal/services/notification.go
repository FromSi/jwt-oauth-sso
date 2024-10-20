package services

type QueryNotificationService interface{}

type MutableNotificationService interface {
	SendText(string)
}

type NotificationService interface {
	QueryNotificationService
	MutableNotificationService
}
