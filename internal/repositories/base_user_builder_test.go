package repositories

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewBaseUserBuilder(t *testing.T) {
	baseUserBuilder := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilder)
}

func TestBaseUserBuilder_New(t *testing.T) {
	baseUserBuilderOne := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilderOne)

	baseUserBuilderOne.SetUUID("1")

	baseUserBuilderTwo := baseUserBuilderOne.New()

	assert.NotNil(t, baseUserBuilderTwo)
	assert.NotEqual(t, baseUserBuilderOne, baseUserBuilderTwo)
}

func TestBaseUserBuilder_NewFromUser(t *testing.T) {
	baseUserBuilderOne := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilderOne)

	baseUserBuilderOne.SetUUID("1")

	user := GormUser{UUID: "2"}

	baseUserBuilderTwo := baseUserBuilderOne.NewFromUser(&user)

	assert.NotNil(t, baseUserBuilderTwo)
	assert.NotEqual(t, baseUserBuilderOne, baseUserBuilderTwo)
}

func TestBaseUserBuilder_Build(t *testing.T) {
	baseUserBuilder := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilder)

	baseUserBuilder.SetUUID("1")

	user, err := baseUserBuilder.Build()

	assert.NoError(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, baseUserBuilder.user.GetUUID(), user.GetUUID())
}

func TestBaseUserBuilder_BuildToGorm(t *testing.T) {
	baseUserBuilder := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilder)

	baseUserBuilder.SetUUID("1")

	user, err := baseUserBuilder.BuildToGorm()

	assert.NoError(t, err)
	assert.NotNil(t, user)

	assert.Equal(t, baseUserBuilder.user.GetUUID(), user.GetUUID())

	baseUserBuilder.SetUUID("")

	user, err = baseUserBuilder.BuildToGorm()

	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestBaseUserBuilder_SetUUID(t *testing.T) {
	baseUserBuilder := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilder)

	baseUserBuilder.SetUUID("1")

	assert.Equal(t, "1", baseUserBuilder.user.GetUUID())

	baseUserBuilder.SetUUID("2")

	assert.Equal(t, "2", baseUserBuilder.user.GetUUID())
}

func TestBaseUserBuilder_SetEmail(t *testing.T) {
	baseUserBuilder := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilder)

	baseUserBuilder.SetEmail("1")

	assert.Equal(t, "1", baseUserBuilder.user.GetEmail())

	baseUserBuilder.SetEmail("2")

	assert.Equal(t, "2", baseUserBuilder.user.GetEmail())
}

func TestBaseUserBuilder_SetPassword(t *testing.T) {
	baseUserBuilder := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilder)

	baseUserBuilder.SetPassword("1")

	assert.Equal(t, "1", baseUserBuilder.user.GetPassword())

	baseUserBuilder.SetPassword("2")

	assert.Equal(t, "2", baseUserBuilder.user.GetPassword())
}

func TestBaseUserBuilder_SetCreatedAt(t *testing.T) {
	baseUserBuilder := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilder)

	baseUserBuilder.SetCreatedAt(1)

	assert.Equal(t, 1, baseUserBuilder.user.GetCreatedAt())

	baseUserBuilder.SetCreatedAt(2)

	assert.Equal(t, 2, baseUserBuilder.user.GetCreatedAt())
}

func TestBaseUserBuilder_SetUpdatedAt(t *testing.T) {
	baseUserBuilder := NewBaseUserBuilder()

	assert.NotNil(t, baseUserBuilder)

	baseUserBuilder.SetUpdatedAt(1)

	assert.Equal(t, 1, baseUserBuilder.user.GetUpdatedAt())

	baseUserBuilder.SetUpdatedAt(2)

	assert.Equal(t, 2, baseUserBuilder.user.GetUpdatedAt())
}
