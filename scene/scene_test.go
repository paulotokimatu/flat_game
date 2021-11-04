package scene_test

import (
	"flat_game"
	"flat_game/entity"
	"flat_game/input"
	"flat_game/scene"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEntity struct {
	name string
	mock.Mock
	entity.BaseEntity
}

func (entity *MockEntity) Name() string {
	return entity.name
}

func (entity *MockEntity) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	entity.Called(game, delta)
}

func TestShouldTickEntities(t *testing.T) {
	delta := float32(1)

	scene := scene.NewScene("foo")

	entity1 := &MockEntity{name: "entity1"}
	entity1.SetPendingRemoval(true)

	entity2 := &MockEntity{name: "entity2"}
	entity2.On("Tick", nil, float32(1)).Return(nil)

	scene.AddEntity(entity1)
	scene.AddEntity(entity2)

	assert.NotNil(t, scene.EntityByName("entity1"))
	assert.NotNil(t, scene.EntityByName("entity2"))

	scene.Tick(nil, delta)

	assert.Nil(t, scene.EntityByName("entity1"))
	assert.NotNil(t, scene.EntityByName("entity2"))

	entity1.AssertNotCalled(t, "Tick", nil, delta)
	entity2.AssertCalled(t, "Tick", nil, delta)
}

type MockKeyEventListener struct {
	mock.Mock
}

func (listener *MockKeyEventListener) OnKeyEvent(key input.Key, event input.KeyEvent) {
	listener.Called(key, event)
}

func TestOnKeyEvent(t *testing.T) {
	scene := scene.Scene{}

	listener := &MockKeyEventListener{}

	scene.AddKeyEventListener(listener)

	listener.On("OnKeyEvent", input.Key(10), input.EventKeyPressed).Return(nil)

	scene.OnKeyEvent(10, input.EventKeyPressed)

	listener.AssertCalled(t, "OnKeyEvent", input.Key(10), input.EventKeyPressed)
}
