package entity_test

import (
	"flat_game"
	"flat_game/entity"
	"flat_game/game"
	"flat_game/internal/graphics"
	"flat_game/utils"
	"testing"

	"github.com/stretchr/testify/mock"
)

type SpriteEntMockGame struct {
	game.Game
	graphics flat_game.IGraphics
}

func (game *SpriteEntMockGame) Graphics() flat_game.IGraphics {
	return game.graphics
}

type SpriteEntMockGraphics struct {
	mock.Mock
	graphics.OpenGl
}

func (graphics *SpriteEntMockGraphics) DrawSprite(texture flat_game.ITexture, position *utils.Vec2, size *utils.Vec2, color *utils.Vec3) {
	graphics.Called(texture, position, size, color)
}

func TestSpriteExtTickShouldDraw(t *testing.T) {
	mockGraphics := SpriteEntMockGraphics{}

	game := SpriteEntMockGame{
		graphics: &mockGraphics,
	}
	ent := entity.NewSpriteEnt(&entity.Config{Name: "foo"}, nil)

	parent := entity.NewBaseEntity(&entity.Config{Name: "parent"})

	mockGraphics.On("DrawSprite", nil, &utils.Vec2{X: 0, Y: 0}, &utils.Vec2{X: 0, Y: 0}, &utils.Vec3{X: 1, Y: 1, Z: 1}).Return(nil)

	ent.Tick(&game, parent, 1.0)

	mockGraphics.AssertCalled(t, "DrawSprite", nil, &utils.Vec2{X: 0, Y: 0}, &utils.Vec2{X: 0, Y: 0}, &utils.Vec3{X: 1, Y: 1, Z: 1})
}