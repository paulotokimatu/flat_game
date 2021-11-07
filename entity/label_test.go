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

type LabelEntMockGame struct {
	game.Game
	graphics flat_game.IGraphics
}

func (game *LabelEntMockGame) Graphics() flat_game.IGraphics {
	return game.graphics
}

type LabelEntMockGraphics struct {
	mock.Mock
	graphics.OpenGl
}

func (graphics *LabelEntMockGraphics) DrawLabel(font flat_game.IFont, text string, position *utils.Vec2, color *utils.Vec3) {
	graphics.Called(font, text, position, color)
}

func TestTickShouldDraw(t *testing.T) {
	mockGraphics := LabelEntMockGraphics{}

	game := LabelEntMockGame{
		graphics: &mockGraphics,
	}
	ent := entity.NewLabelEnt(&entity.Config{Name: "label"}, nil, "foo", nil)

	mockGraphics.On("DrawLabel", nil, "foo", &utils.Vec2{X: 0, Y: 0}, (*utils.Vec3)(nil)).Return(nil)

	ent.Tick(&game, nil, 1.0)

	mockGraphics.AssertCalled(t, "DrawLabel", nil, "foo", &utils.Vec2{X: 0, Y: 0}, (*utils.Vec3)(nil))
}
