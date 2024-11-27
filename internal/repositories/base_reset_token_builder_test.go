package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBaseResetTokenBuilder(t *testing.T) {
	baseResetTokenBuilder := NewBaseResetTokenBuilder()

	assert.NotNil(t, baseResetTokenBuilder)
}

func TestBaseResetTokenBuilder_New(t *testing.T) {
	baseResetTokenBuilderOne := NewBaseResetTokenBuilder()

	assert.NotNil(t, baseResetTokenBuilderOne)

	baseResetTokenBuilderOne.SetToken("1")

	baseResetTokenBuilderTwo := baseResetTokenBuilderOne.New()

	assert.NotNil(t, baseResetTokenBuilderTwo)
	assert.NotEqual(t, baseResetTokenBuilderOne, baseResetTokenBuilderTwo)
}

func TestBaseResetTokenBuilder_NewFromResetToken(t *testing.T) {
	baseResetTokenBuilderOne := NewBaseResetTokenBuilder()

	assert.NotNil(t, baseResetTokenBuilderOne)

	baseResetTokenBuilderOne.SetToken("1")

	user := GormResetToken{Token: "2"}

	baseResetTokenBuilderTwo := baseResetTokenBuilderOne.NewFromResetToken(&user)

	assert.NotNil(t, baseResetTokenBuilderTwo)
	assert.NotEqual(t, baseResetTokenBuilderOne, baseResetTokenBuilderTwo)
}

func TestBaseResetTokenBuilder_Build(t *testing.T) {
	baseResetTokenBuilder := NewBaseResetTokenBuilder()

	assert.NotNil(t, baseResetTokenBuilder)

	baseResetTokenBuilder.SetToken("1")

	resetToken, err := baseResetTokenBuilder.Build()

	assert.NoError(t, err)
	assert.NotNil(t, resetToken)

	assert.Equal(t, baseResetTokenBuilder.resetToken.GetToken(), resetToken.GetToken())
}

func TestBaseResetTokenBuilder_BuildToGorm(t *testing.T) {
	baseResetTokenBuilder := NewBaseResetTokenBuilder()

	assert.NotNil(t, baseResetTokenBuilder)

	baseResetTokenBuilder.SetToken("1")

	resetToken, err := baseResetTokenBuilder.BuildToGorm()

	assert.NoError(t, err)
	assert.NotNil(t, resetToken)

	assert.Equal(t, baseResetTokenBuilder.resetToken.GetToken(), resetToken.GetToken())

	baseResetTokenBuilder.SetToken("")

	resetToken, err = baseResetTokenBuilder.BuildToGorm()

	assert.Error(t, err)
	assert.Nil(t, resetToken)
}

func TestBaseResetTokenBuilder_SetToken(t *testing.T) {
	baseResetTokenBuilder := NewBaseResetTokenBuilder()

	assert.NotNil(t, baseResetTokenBuilder)

	baseResetTokenBuilder.SetToken("1")

	assert.Equal(t, "1", baseResetTokenBuilder.resetToken.GetToken())

	baseResetTokenBuilder.SetToken("2")

	assert.Equal(t, "2", baseResetTokenBuilder.resetToken.GetToken())
}

func TestBaseResetTokenBuilder_SetResetTokenUUID(t *testing.T) {
	baseResetTokenBuilder := NewBaseResetTokenBuilder()

	assert.NotNil(t, baseResetTokenBuilder)

	baseResetTokenBuilder.SetUserUUID("1")

	assert.Equal(t, "1", baseResetTokenBuilder.resetToken.GetUserUUID())

	baseResetTokenBuilder.SetUserUUID("2")

	assert.Equal(t, "2", baseResetTokenBuilder.resetToken.GetUserUUID())
}

func TestBaseResetTokenBuilder_SetExpiresAt(t *testing.T) {
	baseResetTokenBuilder := NewBaseResetTokenBuilder()

	assert.NotNil(t, baseResetTokenBuilder)

	baseResetTokenBuilder.SetExpiresAt(1)

	assert.Equal(t, 1, baseResetTokenBuilder.resetToken.GetExpiresAt())

	baseResetTokenBuilder.SetExpiresAt(2)

	assert.Equal(t, 2, baseResetTokenBuilder.resetToken.GetExpiresAt())
}

func TestBaseResetTokenBuilder_SetCreatedAt(t *testing.T) {
	baseResetTokenBuilder := NewBaseResetTokenBuilder()

	assert.NotNil(t, baseResetTokenBuilder)

	baseResetTokenBuilder.SetCreatedAt(1)

	assert.Equal(t, 1, baseResetTokenBuilder.resetToken.GetCreatedAt())

	baseResetTokenBuilder.SetCreatedAt(2)

	assert.Equal(t, 2, baseResetTokenBuilder.resetToken.GetCreatedAt())
}
