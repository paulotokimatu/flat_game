package game_test

import (
	"testing"

	"github.com/paulotokimatu/flat_game/game"

	"github.com/stretchr/testify/assert"
)

func TestOpenGlGeneration(t *testing.T) {
	lib, err := game.NewGraphics("opengl")

	assert.Nil(t, err, "factory should not return an error")
	assert.Equal(t, lib.Name(), "opengl", "lib type should be opengl")
}

func TestFailsIfInvalidLibName(t *testing.T) {
	_, err := game.NewGraphics("invalid_name")

	expectedError := "invalid graphic library \"invalid_name\""

	assert.NotNil(t, err, "invalid lib should return an error")
	assert.Equal(t, err.Error(), expectedError, "different error than expected")
}
