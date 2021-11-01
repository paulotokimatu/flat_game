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

type LabelExtMockGame struct {
	game.Game
	graphics flat_game.IGraphics
}

func (game *LabelExtMockGame) Graphics() flat_game.IGraphics {
	return game.graphics
}

type LabelExtMockGraphics struct {
	mock.Mock
	graphics.OpenGl
}

func (graphics *LabelExtMockGraphics) DrawLabel(font flat_game.IFont, text string, position *utils.Vec2, color *utils.Vec3) {
	graphics.Called(font, text, position, color)
}

type LabelExtMockEntity struct {
	entity.BaseEntity
}

func TestLabelExtCanAlwaysTick(t *testing.T) {
	ext := ext.NewLabelExt(nil, nil, "foo", nil)

	assert.True(t, ext.CanTick(nil), "canTick should not return false")
}

func TestTickShouldDraw(t *testing.T) {
	mockGraphics := LabelExtMockGraphics{}

	game := LabelExtMockGame{
		graphics: &mockGraphics,
	}
	ext := ext.NewLabelExt(&LabelExtMockEntity{}, nil, "foo", nil)

	mockGraphics.On("DrawLabel", nil, "foo", (*utils.Vec2)(nil), (*utils.Vec3)(nil)).Return(nil)

	ext.Tick(&game, 1.0)

	mockGraphics.AssertCalled(t, "DrawLabel", nil, "foo", (*utils.Vec2)(nil), (*utils.Vec3)(nil))
}
