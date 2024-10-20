package services

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func Test_NewLogNotificationService(t *testing.T) {
	logNotificationService := NewLogNotificationService()

	assert.NotNil(t, logNotificationService)
}

func TestLogNotificationService_SendText(t *testing.T) {
	logNotificationService := NewLogNotificationService()

	buf := bytes.Buffer{}

	log.SetOutput(&buf)
	log.SetFlags(0)

	logNotificationService.SendText("1")

	assert.Equal(t, buf.String(), "1\n")

	buf.Reset()

	logNotificationService.SendText("2")

	assert.Equal(t, buf.String(), "2\n")
}
