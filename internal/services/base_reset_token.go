package services

import (
	"errors"
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

func (receiver *BaseResetTokenService) SendNewResetTokenByUserEmail(email string) error {
	user := receiver.userRepository.GetUserByEmail(email)

	if user == nil {
		return errors.New("user not found")
	}

	newResetToken := repositories.NewGormResetToken()

	newResetToken.SetToken(receiver.GenerateToken())
	newResetToken.SetUserUUID(user.GetUUID())
	newResetToken.SetExpiredAt(int(time.Now().AddDate(0, 0, receiver.config.GetExpirationResetInDays()).Unix()))
	newResetToken.SetCreatedAt(int(time.Now().Unix()))

	err := receiver.resetTokenRepository.CreateResetToken(newResetToken)

	if err != nil {
		return err
	}

	err = receiver.notificationService.SendTextByUser(user, "your reset token is: "+newResetToken.GetToken())

	if err != nil {
		return err
	}

	return nil
}

func (receiver *BaseResetTokenService) ResetPasswordByTokenAndNewPassword(token string, newPassword string) error {
	resetToken := receiver.resetTokenRepository.GetResetTokenByToken(token)

	if resetToken == nil {
		return errors.New("token not found")
	}

	hashedPassword, err := receiver.userService.HashPassword(newPassword)

	if err != nil {
		return err
	}

	err = receiver.userRepository.UpdatePasswordByUUIDAndPasswordAndUpdatedAt(
		resetToken.GetUserUUID(),
		hashedPassword,
		int(time.Now().Unix()),
	)

	if err != nil {
		return err
	}

	err = receiver.resetTokenRepository.DeleteResetToken(token)

	if err != nil {
		return err
	}

	return nil
}

func (receiver *BaseResetTokenService) ResetPasswordByUserUUIDAndOldPasswordAndNewPassword(
	userUUID string,
	oldPassword string,
	newPassword string,
) error {
	userExists := receiver.userRepository.HasUserByUUIDAndPassword(userUUID, oldPassword)

	if !userExists {
		return errors.New("user not found")
	}

	hashedPassword, err := receiver.userService.HashPassword(newPassword)

	if err != nil {
		return err
	}

	err = receiver.userRepository.UpdatePasswordByUUIDAndPasswordAndUpdatedAt(
		userUUID,
		hashedPassword,
		int(time.Now().Unix()),
	)

	if err != nil {
		return err
	}

	return nil
}
