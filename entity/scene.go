package entity

import (
	"flat_game"
	"flat_game/input"
	"flat_game/internal/physics"
)

type Scene struct {
	BaseEntity
	collisions        [][2]flat_game.IEntity
	keyEventListeners []input.IKeyEventListener // It is better to create an object so it only triggers for some keys
}

func NewScene(config *Config) *Scene {
	entity := NewBaseEntity(config)

	return &Scene{
		BaseEntity: *entity,
	}
}

func (scene *Scene) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	physics.ExecuteCollisions(game, scene.collisions)
}

func (scene *Scene) OnKeyEvent(key input.Key, event input.KeyEvent) {
	listeners := scene.keyEventListeners

	for _, listener := range listeners {
		listener.OnKeyEvent(key, event)
	}
}

func (scene *Scene) AddCollision(entityA flat_game.IEntity, entityB flat_game.IEntity) {
	entities := [2]flat_game.IEntity{entityA, entityB}
	scene.collisions = append(scene.collisions, entities)
}

func (scene *Scene) AddKeyEventListener(listener input.IKeyEventListener) {
	scene.keyEventListeners = append(scene.keyEventListeners, listener)
}
