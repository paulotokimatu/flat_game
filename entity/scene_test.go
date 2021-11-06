package entity_test

import (
	"flat_game/entity"
	"flat_game/input"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockKeyEventListener struct {
	mock.Mock
}

func (listener *MockKeyEventListener) OnKeyEvent(key input.Key, event input.KeyEvent) {
	listener.Called(key, event)
}

func TestOnKeyEvent(t *testing.T) {
	scene := entity.Scene{}

	listener := &MockKeyEventListener{}

	scene.AddKeyEventListener(listener)

	listener.On("OnKeyEvent", input.Key(10), input.EventKeyPressed).Return(nil)

	scene.OnKeyEvent(10, input.EventKeyPressed)

	listener.AssertCalled(t, "OnKeyEvent", input.Key(10), input.EventKeyPressed)
}
