package services

import (
	"github.com/fromsi/jwt-oauth-sso/internal/configs"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"time"
)

type BaseResetTokenService struct {
	config               configs.TokenConfig
	userService          UserService
	resetTokenRepository repositories.ResetTokenRepository
	userRepository       repositories.UserRepository
	notificationService  NotificationService
	resetTokenBuilder    repositories.ResetTokenBuilder
}

func NewBaseResetTokenService(
	config configs.TokenConfig,
	userService UserService,
	resetTokenRepository repositories.ResetTokenRepository,
	userRepository repositories.UserRepository,
	notificationService NotificationService,
	resetTokenBuilder repositories.ResetTokenBuilder,
) *BaseResetTokenService {
	return &BaseResetTokenService{
		config:               config,
		userService:          userService,
		resetTokenRepository: resetTokenRepository,
		userRepository:       userRepository,
		notificationService:  notificationService,
		resetTokenBuilder:    resetTokenBuilder,
	}
}

func (receiver *BaseResetTokenService) GenerateToken() string {
	return uuid.New().String()
}

func (receiver *BaseResetTokenService) SendNewResetTokenByUser(user repositories.User) error {
	expiresAt := time.
		Now().
		AddDate(0, 0, receiver.config.GetExpirationResetInDays()).
		Unix()

	newResetToken, err := receiver.resetTokenBuilder.
		New().
		SetToken(receiver.GenerateToken()).
		SetUserUUID(user.GetUUID()).
		SetExpiresAt(int(expiresAt)).
		SetCreatedAt(int(time.Now().Unix())).
		Build()

	if err != nil {
		return err
	}

	err = receiver.resetTokenRepository.CreateResetToken(newResetToken)

	if err != nil {
		return err
	}

	err = receiver.notificationService.SendTextByUser(user, "your reset token: "+newResetToken.GetToken())

	if err != nil {
		return err
	}

	return nil
}
