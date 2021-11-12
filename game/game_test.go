package game_test

import (
	"testing"

	"github.com/paulotokimatu/flat_game"

	"github.com/paulotokimatu/flat_game/entity"
	"github.com/paulotokimatu/flat_game/game"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type GameMockGraphics struct {
	flat_game.IGraphics
}

func (entity *GameMockGraphics) Setup(config *flat_game.Config, onKeyEvent flat_game.OnKeyEventFunction) {
}

func (entity *GameMockGraphics) Tick() {}

type GameMockEntity struct {
	name string
	mock.Mock
	entity.BaseEntity
}

func (entity *GameMockEntity) Name() string {
	return entity.name
}

func (entity *GameMockEntity) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	entity.Called(game, parent, delta)
}

func TestShouldTickEntities(t *testing.T) {
	mockGame := game.NewGameWithGraphics(
		flat_game.Config{},
		&GameMockGraphics{},
	)

	delta := float32(1)

	scene := entity.NewScene(&entity.Config{Name: "foo"})
	mockGame.SetScene(scene, false)

	entity1 := &GameMockEntity{name: "entity1"}
	entity1.SetPendingRemoval(true)

	entity2 := &GameMockEntity{name: "entity2"}
	entity2.On("Tick", mockGame, scene, delta).Return(nil)

	scene.AddChild(entity1)
	scene.AddChild(entity2)

	mockGame.Tick(delta)

	assert.Nil(t, scene.ChildByName("entity1"))
	assert.NotNil(t, scene.ChildByName("entity2"))

	entity1.AssertNotCalled(t, "Tick", mockGame, scene, delta)
	entity2.AssertCalled(t, "Tick", mockGame, scene, delta)
}
