package flat_game

import (
	"flat_game/input"
	"flat_game/utils"
)

type IGraphics interface {
	DrawSprite(texture ITexture, position *utils.Vec2, size *utils.Vec2, color utils.Vec3)

	IsKeyPressed(key input.Key) bool

	PostTick()

	PreTick()

	// Setup(GameConfig, IGame)
	Setup(config *GameConfig)

	Terminate()

	Tick()

	WindowShouldClose() bool
}
