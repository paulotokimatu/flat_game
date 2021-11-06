package flat_game

import (
	"flat_game/utils"
)

type IEntity interface {
	Name() string

	OnCollision(externalEntity IEntity)

	Position() *utils.Vec2

	SetPosition(position *utils.Vec2)

	Size() *utils.Vec2

	IsPendingRemoval() bool

	SetPendingRemoval(pendingRemoval bool)

	// Ticker

	AddChild(entity IEntity)

	CommitChild(childToAdd IEntity)

	CanTick(game IGame) bool

	ChildrenNames() []string

	ChildByName(name string) IEntity

	ChildrenToAdd() []IEntity

	ClearChildrenToAdd()

	RemoveChild(child IEntity)

	Tick(game IGame, parent IEntity, delta float32)

	UpdateChildrenNames(childrenNames []string)
}
