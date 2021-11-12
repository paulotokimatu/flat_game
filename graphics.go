package flat_game

import (
	"github.com/paulotokimatu/flat_game/input"
	"github.com/paulotokimatu/flat_game/utils"
)

type OnKeyEventFunction func(key input.Key, event input.KeyEvent)

type IGraphics interface {
	DrawLabel(font IFont, text string, position *utils.Vec2, color *utils.Vec3)

	DrawSprite(texture ITexture, position *utils.Vec2, size *utils.Vec2, color *utils.Vec3)

	IsKeyPressed(key input.Key) bool

	Name() string

	PostTick()

	PreTick()

	Setup(config *Config, onKeyEvent OnKeyEventFunction)

	Terminate()

	Tick()

	WindowShouldClose() bool
}
