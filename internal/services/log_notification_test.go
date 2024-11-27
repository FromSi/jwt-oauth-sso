package services

import (
	"bytes"
	repositories_mocks "github.com/fromsi/jwt-oauth-sso/mocks/repositories"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log"
	"testing"
)

func Test_NewLogNotificationService(t *testing.T) {
	logNotificationService := NewLogNotificationService()

	assert.NotNil(t, logNotificationService)
}

func TestLogNotificationService_SendTextByUser(t *testing.T) {
	logNotificationService := NewLogNotificationService()

	buf := bytes.Buffer{}

	log.SetOutput(&buf)
	log.SetFlags(0)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockUser := repositories_mocks.NewMockUser(mockController)

	mockUser.EXPECT().GetUUID().Return("1")

	err := logNotificationService.SendTextByUser(mockUser, "1")

	assert.NoError(t, err)

	assert.Equal(t, buf.String(), "user: 1 | text: 1\n")

	buf.Reset()

	mockUser.EXPECT().GetUUID().Return("2")

	err = logNotificationService.SendTextByUser(mockUser, "2")

	assert.NoError(t, err)

	assert.Equal(t, buf.String(), "user: 2 | text: 2\n")
}
