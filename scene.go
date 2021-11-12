package flat_game

import "github.com/paulotokimatu/flat_game/input"

type IScene interface {
	IEntity

	AddCollision(entityA IEntity, entityB IEntity)

	AddKeyEventListener(listener input.IKeyEventListener)

	OnKeyEvent(key input.Key, event input.KeyEvent)
}
