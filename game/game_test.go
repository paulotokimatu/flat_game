package game_test

import (
	"flat_game"
	"flat_game/entity"
	"flat_game/game"
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
	entity.Called(game, parent, delta)
}

func TestShouldTickEntities(t *testing.T) {
	mockGame := &game.Game{}

	delta := float32(1)

	scene := entity.NewScene(&entity.Config{Name: "foo"})
	mockGame.SetScene(scene)

	entity1 := &MockEntity{name: "entity1"}
	entity1.SetPendingRemoval(true)

	entity2 := &MockEntity{name: "entity2"}
	entity2.On("Tick", mockGame, scene, delta).Return(nil)

	scene.AddChild(entity1)
	scene.AddChild(entity2)

	mockGame.Tick(delta)

	assert.Nil(t, scene.ChildByName("entity1"))
	assert.NotNil(t, scene.ChildByName("entity2"))

	entity1.AssertNotCalled(t, "Tick", mockGame, scene, delta)
	entity2.AssertCalled(t, "Tick", mockGame, scene, delta)
}
