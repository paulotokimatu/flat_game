package game

import (
	"flat_game"
	"flat_game/input"
	graphicsLib "flat_game/internal/graphics"
	"fmt"
	"time"
)

type Game struct {
	config       flat_game.Config
	currentScene flat_game.IScene
	fonts        map[string]flat_game.IFont
	graphics     flat_game.IGraphics
	lastTick     time.Time
	scenes       map[string]flat_game.IScene
	textures     map[string]flat_game.ITexture
}

func NewGame(config flat_game.Config) *Game {
	graphics, err := NewGraphics(GraphicsLib(config.Graphics))
	if err != nil {
		panic(err)
	}

	return NewGameWithGraphics(config, graphics)
}

func NewGameWithGraphics(config flat_game.Config, graphics flat_game.IGraphics) *Game {
	game := &Game{
		config:   config,
		fonts:    map[string]flat_game.IFont{},
		graphics: graphics,
		scenes:   map[string]flat_game.IScene{},
		textures: map[string]flat_game.ITexture{},
	}

	game.graphics.Setup(&config, game.onKeyEvent)

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

		game.Tick(delta)

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

func (game *Game) Tick(delta float32) {
	game.graphics.Tick()

	runTick(game, nil, game.currentScene, delta)
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

func (game *Game) AddScene(scene flat_game.IScene) {
	if game.scenes == nil {
		game.scenes = map[string]flat_game.IScene{}
	}

	game.scenes[scene.Name()] = scene
}

func (game *Game) CurrentScene() flat_game.IScene {
	return game.currentScene
}

func (game *Game) DeleteScene(sceneName string) {
	if game.CurrentScene().Name() == sceneName {
		fmt.Println(sceneName)
		panic("cannot delete current scene")
	}

	delete(game.scenes, sceneName)
}

func (game *Game) SceneByName(sceneName string) flat_game.IScene {
	return game.scenes[sceneName]
}

func (game *Game) SetScene(scene flat_game.IScene, deletePreviousScene bool) {
	previousScene := game.CurrentScene()

	if _, ok := game.scenes[scene.Name()]; !ok {
		game.AddScene(scene)
	}

	game.currentScene = scene

	if game.currentScene != nil && deletePreviousScene {
		game.DeleteScene(previousScene.Name())
	}
}

func (game *Game) AddFont(
	name string,
	fileName string,
	charCodeMin rune,
	charCodeMax rune,
	scale int,
) (flat_game.IFont, error) {
	font, err := graphicsLib.NewFontFromFile(name, fileName, charCodeMin, charCodeMax, scale)
	if err != nil {
		return nil, err
	}

	game.fonts[name] = font

	return font, nil
}

func (game *Game) FontByName(name string) flat_game.IFont {
	return game.fonts[name]
}

func runTick(game flat_game.IGame, parent flat_game.IEntity, entity flat_game.IEntity, delta float32) {
	for _, childToAdd := range entity.ChildrenToAdd() {
		entity.CommitChild(childToAdd)
	}
	entity.ClearChildrenToAdd()

	childrenToPersist := []string{}

	for _, childName := range entity.ChildrenNames() {
		child := entity.ChildByName(childName)

		if child.IsPendingRemoval() {
			entity.RemoveChild(child)

			continue
		}

		childrenToPersist = append(childrenToPersist, child.Name())

		runTick(game, entity, child, delta)
	}

	entity.Tick(game, parent, delta)

	entity.UpdateChildrenNames(childrenToPersist)
}

func (game *Game) onKeyEvent(key input.Key, event input.KeyEvent) {
	game.currentScene.OnKeyEvent(key, event)
}
