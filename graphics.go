package flat_game

import (
	"flat_game/input"
	"flat_game/utils"
)

type IGraphics interface {
	DrawLabel(font IFont, text string, position *utils.Vec2, color utils.Vec3)

	DrawSprite(texture ITexture, position *utils.Vec2, size *utils.Vec2, color utils.Vec3)

	IsKeyPressed(key input.Key) bool

	Name() string

	PostTick()

	PreTick()

	// Setup(Config, IGame)
	Setup(config *Config)

	Terminate()

	Tick()

	WindowShouldClose() bool
}
