package entity

import (
	"flat_game"
	"flat_game/utils"
)

type BaseEntity struct {
	keyEventListeners []flat_game.IExt
	name              string
	exts              []flat_game.IExt
	pendingRemoval    bool
	position          *utils.Vec2
	size              *utils.Vec2
}

func NewEntity(Config *Config) *BaseEntity {
	position := utils.Vec2{
		X: Config.Position.X,
		Y: Config.Position.Y,
	}
	size := utils.Vec2{
		X: Config.Size.X,
		Y: Config.Size.Y,
	}
	entity := &BaseEntity{
		keyEventListeners: nil,
		name:              Config.Name,
		exts:              nil,
		pendingRemoval:    false,
		position:          &position,
		size:              &size,
	}

	return entity
}

func (entity *BaseEntity) Tick(game flat_game.IGame, delta float32) {
	for i := 0; i < len(entity.exts); i++ {
		if entity.exts[i].CanTick(game) {
			entity.exts[i].Tick(game, delta)
		}
	}
}

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

func (entity *BaseEntity) AddExt(node flat_game.IExt) {
	entity.exts = append(entity.exts, node)
}

func (entity *BaseEntity) OnCollision(externalEntity flat_game.IEntity) {
}

func (entity *BaseEntity) IsPendingRemoval() bool {
	return entity.pendingRemoval
}

func (entity *BaseEntity) SetPendingRemoval(pendingRemoval bool) {
	entity.pendingRemoval = pendingRemoval
}
