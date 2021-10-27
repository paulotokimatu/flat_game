package ext_test

import (
	"flat_game"
	"flat_game/entity"
	"flat_game/ext"
	"flat_game/game"
	"flat_game/internal/graphics"
	"flat_game/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGame struct {
	game.Game
	graphics flat_game.IGraphics
}

func (game *MockGame) Graphics() flat_game.IGraphics {
	return game.graphics
}

type MockGraphics struct {
	mock.Mock
	graphics.OpenGl
}

func (graphics *MockGraphics) DrawSprite(texture flat_game.ITexture, position *utils.Vec2, size *utils.Vec2, color utils.Vec3) {
	graphics.Called(texture, position, size, color)
}

type MockEntity struct {
	entity.BaseEntity
}

func (entity *MockEntity) Position() *utils.Vec2 {
	return nil
}

func (entity *MockEntity) Size() *utils.Vec2 {
	return nil
}

func TestCanAlwaysTick(t *testing.T) {
	ext := ext.NewSpriteExt(nil, nil)

	assert.True(t, ext.CanTick(nil), "canTick should not return false")
}

func TestTickShouldDraw(t *testing.T) {
	mockGraphics := MockGraphics{}

	game := MockGame{
		graphics: &mockGraphics,
	}
	ext := ext.NewSpriteExt(&MockEntity{}, nil)

	mockGraphics.On("DrawSprite", nil, (*utils.Vec2)(nil), (*utils.Vec2)(nil), utils.Vec3{X: 1, Y: 1, Z: 1}).Return(nil)

	ext.Tick(&game, 1.0)

	mockGraphics.AssertCalled(t, "DrawSprite", nil, (*utils.Vec2)(nil), (*utils.Vec2)(nil), utils.Vec3{X: 1, Y: 1, Z: 1})
}
