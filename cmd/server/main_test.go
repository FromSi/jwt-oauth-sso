package main

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
	"testing"
)

func Test_CreateApp(t *testing.T) {
	assert.Nil(t, fx.ValidateApp(CreateApp()))
}
