package flat_game

import (
	"flat_game/input"
)

type IGame interface {
	AddCollision(entityA IEntity, entityB IEntity)

	AddEntity(entity IEntity)

	AddKeyEventListener(listener input.IKeyEventListener)

	AddTexture(name string, fileName string) (ITexture, error)

	Config() GameConfig

	EntityByName(name string) IEntity

	Graphics() IGraphics

	OnKeyEvent(key input.Key, event input.KeyEvent)

	TextureByName(name string) ITexture
}
