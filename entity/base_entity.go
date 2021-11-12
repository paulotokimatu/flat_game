package entity

import (
	"github.com/paulotokimatu/flat_game"

	"github.com/paulotokimatu/flat_game/utils"
)

type BaseEntity struct {
	children          map[string]flat_game.IEntity
	childrenNames     []string
	childrenToAdd     []flat_game.IEntity
	customData        map[string]interface{}
	keyEventListeners []flat_game.IEntity
	name              string
	pendingRemoval    bool
	position          *utils.Vec2
	size              *utils.Vec2
}

func NewBaseEntity(config *Config) *BaseEntity {
	children := map[string]flat_game.IEntity{}
	childrenNames := []string{}

	for _, child := range config.Children {
		childrenNames = append(childrenNames, child.Name())
		children[child.Name()] = child
	}

	entity := &BaseEntity{
		keyEventListeners: nil,
		name:              config.Name,
		children:          children,
		childrenNames:     childrenNames,
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
	entity.childrenToAdd = append(entity.childrenToAdd, child)
}

func (entity *BaseEntity) CommitChild(childToAdd flat_game.IEntity) {
	entity.childrenNames = append(entity.childrenNames, childToAdd.Name())
	entity.children[childToAdd.Name()] = childToAdd
}

func (entity *BaseEntity) ChildByName(name string) flat_game.IEntity {
	return entity.children[name]
}

func (entity *BaseEntity) ChildrenNames() []string {
	return entity.childrenNames
}

func (entity *BaseEntity) UpdateChildrenNames(childrenNames []string) {
	entity.childrenNames = childrenNames
}

func (entity *BaseEntity) ChildrenToAdd() []flat_game.IEntity {
	return entity.childrenToAdd
}

func (entity *BaseEntity) ClearChildrenToAdd() {
	entity.childrenToAdd = []flat_game.IEntity{}
}

func (entity *BaseEntity) RemoveChild(child flat_game.IEntity) {
	delete(entity.children, child.Name())
}

func (entity *BaseEntity) OnCollision(game flat_game.IGame, externalEntity flat_game.IEntity) {}

func (entity *BaseEntity) IsPendingRemoval() bool {
	return entity.pendingRemoval
}

func (entity *BaseEntity) SetPendingRemoval(pendingRemoval bool) {
	entity.pendingRemoval = pendingRemoval
}

func (entity *BaseEntity) CustomData(key string) interface{} {
	return entity.customData[key]
}

func (entity *BaseEntity) SetCustomData(key string, data interface{}) {
	entity.customData[key] = data
}
