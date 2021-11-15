package entity_test

import (
	"testing"

	"github.com/paulotokimatu/flat_game"
	"github.com/paulotokimatu/flat_game/entity"
	"github.com/paulotokimatu/flat_game/game"
	"github.com/paulotokimatu/flat_game/internal/graphics"
	"github.com/paulotokimatu/flat_game/utils"

	"github.com/stretchr/testify/mock"
)

type LabelEntDoubleGame struct {
	game.Game
	graphics flat_game.IGraphics
}

func (game *LabelEntDoubleGame) Graphics() flat_game.IGraphics {
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

	game := LabelEntDoubleGame{
		graphics: &mockGraphics,
	}
	ent := entity.NewLabelEnt(&entity.Config{Name: "label"}, nil, "foo", nil)

	mockGraphics.On("DrawLabel", nil, "foo", &utils.Vec2{X: 0, Y: 0}, (*utils.Vec3)(nil)).Return(nil)

	ent.Tick(&game, nil, 1.0)

	mockGraphics.AssertCalled(t, "DrawLabel", nil, "foo", &utils.Vec2{X: 0, Y: 0}, (*utils.Vec3)(nil))
}
