package flat_game

import (
	"flat_game/utils"
)

type IEntity interface {
	AddExt(node IExt)

	Name() string

	OnCollision(externalEntity IEntity)

	Position() *utils.Vec2

	SetPosition(position *utils.Vec2)

	Size() *utils.Vec2

	Tick(game IGame, delta float32)

	IsPendingRemoval() bool

	SetPendingRemoval(pendingRemoval bool)
}
