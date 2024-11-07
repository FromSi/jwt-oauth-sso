package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"log"
)

type LogNotificationService struct{}

func NewLogNotificationService() *LogNotificationService {
	return &LogNotificationService{}
}

func (receiver *LogNotificationService) SendTextByUser(user repositories.User, text string) error {
	log.Println("user", user.GetUUID(), "text", text)

	return nil
}
