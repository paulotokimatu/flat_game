package game

import (
	"flat_game"
	"flat_game/input"
	graphicsLib "flat_game/internal/graphics"
	"flat_game/internal/physics"
	"time"
)

type Game struct {
	collisions        [][2]flat_game.IEntity
	config            flat_game.GameConfig
	entities          map[string]flat_game.IEntity
	graphics          flat_game.IGraphics
	keyEventListeners []input.IKeyEventListener // It is better to create an object so it only triggers for some keys
	lastTick          time.Time
	textures          map[string]flat_game.ITexture
}

func NewGame(config flat_game.GameConfig) *Game {
	// TODO think about injecting this factory
	graphics := NewGraphics(GraphicsLib(config.Graphics))

	game := &Game{
		config:   config,
		entities: map[string]flat_game.IEntity{},
		graphics: graphics,
		textures: map[string]flat_game.ITexture{},
	}

	// game.graphics.Setup(config, game)
	game.graphics.Setup(&config)

	return game
}

func (game *Game) Run() {
	defer game.graphics.Terminate()

	game.lastTick = time.Now()

	for !game.graphics.WindowShouldClose() {
		timeStartTick := time.Now()
		delta := float32(time.Since(game.lastTick).Seconds())
		game.lastTick = timeStartTick

		game.preTick()

		game.tick(delta)

		game.postTick()

		time.Sleep(time.Second/time.Duration(game.config.MaxFPS) - time.Since(timeStartTick))
	}
}

// TODO Maybe it is better to just create a new instance of Game...
func (game *Game) Stop() {
	game.collisions = nil
	game.keyEventListeners = nil
	game.entities = nil
}

func (game *Game) OnKeyEvent(key input.Key, event input.KeyEvent) {
	for i := 0; i < len(game.keyEventListeners); i++ {
		game.keyEventListeners[i].OnKeyEvent(key, event)
	}
}

func (game *Game) SetGraphicsLib(graphics flat_game.IGraphics) {
	game.graphics = graphics
}

func (game *Game) Graphics() flat_game.IGraphics {
	return game.graphics
}

func (game *Game) Config() flat_game.GameConfig {
	return game.config
}

func (game *Game) Terminate() {
}

func (game *Game) preTick() {
	game.graphics.PreTick()
}

func (game *Game) tick(delta float32) {
	game.graphics.Tick()

	for _, entity := range game.entities {
		// Maybe use a state machine instead
		if entity.IsPendingRemoval() {
			game.removeEntity(entity)
		} else {
			entity.Tick(delta)
		}
	}

	physics.ExecuteCollisions(game.collisions)
}

func (game *Game) postTick() {
	game.graphics.PostTick()
}

func (game *Game) AddEntity(entity flat_game.IEntity) {
	game.entities[entity.Name()] = entity
}

func (game *Game) EntityByName(name string) flat_game.IEntity {
	return game.entities[name]
}

func (game *Game) removeEntity(entity flat_game.IEntity) {
	delete(game.entities, entity.Name())
}

func (game *Game) AddCollision(entityA flat_game.IEntity, entityB flat_game.IEntity) {
	entities := [2]flat_game.IEntity{entityA, entityB}
	game.collisions = append(game.collisions, entities)
}

func (game *Game) AddKeyEventListener(listener input.IKeyEventListener) {
	game.keyEventListeners = append(game.keyEventListeners, listener)
}

func (game *Game) AddTexture(name string, fileName string) (flat_game.ITexture, error) {
	texture, err := graphicsLib.NewTextureFromFile(name, fileName)
	if err != nil {
		return texture, err
	}

	game.textures[texture.Name] = texture

	return texture, nil
}

func (game *Game) TextureByName(name string) flat_game.ITexture {
	return game.textures[name]
}
