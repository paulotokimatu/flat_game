package entity_test

import (
	"flat_game"
	"flat_game/entity"
	"flat_game/utils"
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockExt struct {
	canTick bool
	mock.Mock
}

func (ext *MockExt) CanTick(game flat_game.IGame) bool {
	return ext.canTick
}

func (ext *MockExt) Tick(game flat_game.IGame, delta float32) {
	ext.Called(game, delta)
}

func TestTickShouldCallTicksOfExts(t *testing.T) {
	delta := float32(1)

	entityConfig := entity.EntityConfig{
		Name:     "square",
		Position: utils.Vec2{X: 0, Y: 0},
		Size:     utils.Vec2{X: 10, Y: 10},
	}
	entity := entity.NewEntity(&entityConfig)

	ext1 := &MockExt{canTick: false}
	ext2 := &MockExt{canTick: true}

	ext2.On("Tick", nil, delta).Return(nil)

	entity.AddExt(ext1)
	entity.AddExt(ext2)

	entity.Tick(nil, delta)

	ext1.AssertNotCalled(t, "Tick", nil, delta)
	ext2.AssertCalled(t, "Tick", nil, delta)
}
