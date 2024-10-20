package services

import "log"

type LogNotificationService struct{}

func NewLogNotificationService() *LogNotificationService {
	return &LogNotificationService{}
}

func (receiver LogNotificationService) SendText(text string) {
	log.Println(text)
}
