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
	name              string
}

func NewScene(name string) *Scene {
	return &Scene{
		entities: map[string]flat_game.IEntity{},
		name:     name,
	}
}

func (scene *Scene) Name() string {
	return scene.name
}

func (scene *Scene) Tick(game flat_game.IGame, delta float32) {
	for _, entity := range scene.entities {
		// I have to treat this as a special case since this entity has no parent
		if entity.IsPendingRemoval() {
			scene.RemoveChild(entity)

			continue
		}

		startTick(game, nil, entity, delta)
	}

	physics.ExecuteCollisions(scene.collisions)
}

func (scene *Scene) AddEntity(entity flat_game.IEntity) {
	scene.entities[entity.Name()] = entity
}

func (scene *Scene) EntityByName(name string) flat_game.IEntity {
	return scene.entities[name]
}

func (scene *Scene) RemoveChild(entity flat_game.IEntity) {
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

func startTick(game flat_game.IGame, parent flat_game.IEntity, entity flat_game.IEntity, delta float32) {
	if entity.IsPendingRemoval() {
		parent.RemoveChild(entity)

		return
	}

	if !entity.CanTick(game) {
		return
	}

	for _, child := range entity.Children() {
		startTick(game, entity, child, delta)
	}

	entity.Tick(game, parent, delta)
}
