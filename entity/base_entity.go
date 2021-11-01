package entity

import (
	"flat_game"
	"flat_game/utils"
)

type BaseEntity struct {
	children          map[string]flat_game.IEntity
	keyEventListeners []flat_game.IExt
	name              string
	pendingRemoval    bool
	position          *utils.Vec2
	size              *utils.Vec2
}

func NewBaseEntity(config *Config) *BaseEntity {
	entity := &BaseEntity{
		keyEventListeners: nil,
		name:              config.Name,
		children:          map[string]flat_game.IEntity{},
		pendingRemoval:    false,
		position:          &config.Position,
		size:              &config.Size,
	}

	return entity
}

func (entity *BaseEntity) CanTick(game flat_game.IGame) bool {
	return true
}

func (entity *BaseEntity) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {}

func (entity *BaseEntity) Name() string {
	return entity.name
}

func (entity *BaseEntity) Position() *utils.Vec2 {
	return entity.position
}

func (entity *BaseEntity) SetPosition(position *utils.Vec2) {
	entity.position = position
}

func (entity *BaseEntity) Size() *utils.Vec2 {
	return entity.size
}

func (entity *BaseEntity) SetSize(size *utils.Vec2) {
	entity.size = size
}

func (entity *BaseEntity) AddChild(child flat_game.IEntity) {
	entity.children[child.Name()] = child
}

func (entity *BaseEntity) Children() map[string]flat_game.IEntity {
	return entity.children
}

func (entity *BaseEntity) RemoveChild(child flat_game.IEntity) {
	delete(entity.children, child.Name())
}

func (entity *BaseEntity) OnCollision(externalEntity flat_game.IEntity) {
}

func (entity *BaseEntity) IsPendingRemoval() bool {
	return entity.pendingRemoval
}

func (entity *BaseEntity) SetPendingRemoval(pendingRemoval bool) {
	entity.pendingRemoval = pendingRemoval
}
