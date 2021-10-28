package game

import (
	"flat_game"
	graphicsLib "flat_game/internal/graphics"
	"time"
)

type Game struct {
	config       flat_game.Config
	currentScene flat_game.IScene
	graphics     flat_game.IGraphics
	lastTick     time.Time
	textures     map[string]flat_game.ITexture
}

func NewGame(config flat_game.Config) *Game {
	// TODO think about injecting this factory
	graphics, err := NewGraphics(GraphicsLib(config.Graphics))
	if err != nil {
		panic(err)
	}

	game := &Game{
		config:   config,
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

func (game *Game) Stop() {
	game.currentScene = nil
}

func (game *Game) SetGraphicsLib(graphics flat_game.IGraphics) {
	game.graphics = graphics
}

func (game *Game) Graphics() flat_game.IGraphics {
	return game.graphics
}

func (game *Game) Config() flat_game.Config {
	return game.config
}

func (game *Game) Terminate() {
}

func (game *Game) preTick() {
	game.graphics.PreTick()
}

func (game *Game) tick(delta float32) {
	game.graphics.Tick()

	game.currentScene.Tick(game, delta)
}

func (game *Game) postTick() {
	game.graphics.PostTick()
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

func (game *Game) Scene() flat_game.IScene {
	return game.currentScene
}

func (game *Game) SetScene(scene flat_game.IScene) {
	game.currentScene = scene
}
