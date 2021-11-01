package flat_game

import (
	"flat_game/utils"
)

type IEntity interface {
	AddChild(entity IEntity)

	CanTick(game IGame) bool

	Children() map[string]IEntity

	Name() string

	OnCollision(externalEntity IEntity)

	Position() *utils.Vec2

	RemoveChild(child IEntity)

	SetPosition(position *utils.Vec2)

	Size() *utils.Vec2

	Tick(game IGame, parent IEntity, delta float32)

	IsPendingRemoval() bool

	SetPendingRemoval(pendingRemoval bool)
}
