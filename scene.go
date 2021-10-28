package flat_game

import "flat_game/input"

type IScene interface {
	AddCollision(entityA IEntity, entityB IEntity)

	AddEntity(entity IEntity)

	AddKeyEventListener(listener input.IKeyEventListener)

	EntityByName(name string) IEntity

	OnKeyEvent(key input.Key, event input.KeyEvent)

	Tick(game IGame, delta float32)
}
