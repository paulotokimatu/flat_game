package scene

import (
	"flat_game"
	"flat_game/input"
	"flat_game/internal/physics"
)

type Scene struct {
	collisions        [][2]flat_game.IEntity
	entities          map[string]flat_game.IEntity
	keyEventListeners []input.IKeyEventListener // It is better to create an object so it only triggers for some keys
}

func NewScene() *Scene {
	return &Scene{
		entities: map[string]flat_game.IEntity{},
	}
}

func (scene *Scene) Tick(game flat_game.IGame, delta float32) {
	for _, entity := range scene.entities {
		// Maybe use a state machine instead
		if entity.IsPendingRemoval() {
			scene.removeEntity(entity)
		} else {
			entity.Tick(game, delta)
		}
	}

	physics.ExecuteCollisions(scene.collisions)
}

func (scene *Scene) AddEntity(entity flat_game.IEntity) {
	scene.entities[entity.Name()] = entity
}

func (scene *Scene) EntityByName(name string) flat_game.IEntity {
	return scene.entities[name]
}

func (scene *Scene) removeEntity(entity flat_game.IEntity) {
	delete(scene.entities, entity.Name())
}

func (scene *Scene) OnKeyEvent(key input.Key, event input.KeyEvent) {
	for i := 0; i < len(scene.keyEventListeners); i++ {
		scene.keyEventListeners[i].OnKeyEvent(key, event)
	}
}

func (scene *Scene) AddCollision(entityA flat_game.IEntity, entityB flat_game.IEntity) {
	entities := [2]flat_game.IEntity{entityA, entityB}
	scene.collisions = append(scene.collisions, entities)
}

func (scene *Scene) AddKeyEventListener(listener input.IKeyEventListener) {
	scene.keyEventListeners = append(scene.keyEventListeners, listener)
}
