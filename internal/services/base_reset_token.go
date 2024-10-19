package services

import "github.com/google/uuid"

type BaseResetTokenService struct{}

func NewBaseResetTokenService() *BaseResetTokenService {
	return &BaseResetTokenService{}
}

func (receiver BaseResetTokenService) GenerateToken() string {
	return uuid.New().String()
}
