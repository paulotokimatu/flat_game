package game_test

import (
	"flat_game/game"
	"testing"
)

func TestOpenGlGeneration(t *testing.T) {
	lib, err := game.NewGraphics("opengl")

	if err != nil {
		t.Fatalf("factory should not return an error")
	}

	if lib.Name() != "opengl" {
		t.Fatalf("lib type should be opengl")
	}
}

func TestFailsIfInvalidLibName(t *testing.T) {
	_, err := game.NewGraphics("invalid_name")

	expectedError := "invalid graphic library \"invalid_name\""

	if err != nil && err.Error() != expectedError {
		t.Fatalf("%s should return an error", "invalid_name")
	}
}
