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
}

func NewBaseResetTokenService(
	config configs.TokenConfig,
	userService UserService,
	resetTokenRepository repositories.ResetTokenRepository,
	userRepository repositories.UserRepository,
	notificationService NotificationService,
) *BaseResetTokenService {
	return &BaseResetTokenService{
		config:               config,
		userService:          userService,
		resetTokenRepository: resetTokenRepository,
		userRepository:       userRepository,
		notificationService:  notificationService,
	}
}

func (receiver *BaseResetTokenService) GenerateToken() string {
	return uuid.New().String()
}

func (receiver *BaseResetTokenService) SendNewResetTokenByUser(user repositories.User) error {
	newResetToken := repositories.NewGormResetToken()

	newResetToken.SetToken(receiver.GenerateToken())
	newResetToken.SetUserUUID(user.GetUUID())
	newResetToken.SetExpiresAt(int(time.Now().AddDate(0, 0, receiver.config.GetExpirationResetInDays()).Unix()))
	newResetToken.SetCreatedAt(int(time.Now().Unix()))

	err := receiver.resetTokenRepository.CreateResetToken(newResetToken)

	if err != nil {
		return err
	}

	err = receiver.notificationService.SendTextByUser(user, "your reset token: "+newResetToken.GetToken())

	if err != nil {
		return err
	}

	return nil
}
