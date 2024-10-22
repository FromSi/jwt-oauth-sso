package services

import (
	"errors"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/google/uuid"
	"time"
)

type BaseResetTokenService struct {
	userService          UserService
	resetTokenRepository repositories.ResetTokenRepository
	userRepository       repositories.UserRepository
}

func NewBaseResetTokenService(
	userService UserService,
	resetTokenRepository repositories.ResetTokenRepository,
	userRepository repositories.UserRepository,
) *BaseResetTokenService {
	return &BaseResetTokenService{
		userService:          userService,
		resetTokenRepository: resetTokenRepository,
		userRepository:       userRepository,
	}
}

func (receiver *BaseResetTokenService) GenerateToken() string {
	return uuid.New().String()
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

	err = receiver.userRepository.UpdatePassword(resetToken.GetUserUUID(), hashedPassword, int(time.Now().Unix()))

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

	err = receiver.userRepository.UpdatePassword(userUUID, hashedPassword, int(time.Now().Unix()))

	if err != nil {
		return err
	}

	return nil
}
