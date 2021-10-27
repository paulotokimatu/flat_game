package game_test

import (
	"flat_game/game"
	"flat_game/input"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestRun(t *testing.T) {
	// TODO...
}

type MockKeyEventListener struct {
	mock.Mock
}

func (listener *MockKeyEventListener) OnKeyEvent(key input.Key, event input.KeyEvent) {
	listener.Called(key, event)
}

func TestOnKeyEvent(t *testing.T) {
	game := game.Game{}

	listener := &MockKeyEventListener{}

	game.AddKeyEventListener(listener)

	listener.On("OnKeyEvent", input.Key(10), input.EventKeyPressed).Return(nil)

	game.OnKeyEvent(10, input.EventKeyPressed)

	listener.AssertCalled(t, "OnKeyEvent", input.Key(10), input.EventKeyPressed)
}
