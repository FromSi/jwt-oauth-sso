package services

import (
	"bytes"
	"github.com/fromsi/jwt-oauth-sso/internal/repositories"
	"github.com/stretchr/testify/assert"
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

	user := repositories.NewGormUser()
	user.SetUUID("1")

	err := logNotificationService.SendTextByUser(user, "1")

	assert.Nil(t, err)

	assert.Equal(t, buf.String(), "user 1 text 1\n")
}
