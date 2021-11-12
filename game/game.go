package game

import (
	"time"

	"github.com/paulotokimatu/flat_game"

	"github.com/paulotokimatu/flat_game/input"
)

type Game struct {
	AssetsManager
	ScenesManager
	config   flat_game.Config
	graphics flat_game.IGraphics
	lastTick time.Time
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
		AssetsManager: *NewAssetsManager(),
		ScenesManager: *NewScenesManager(),
		config:        config,
		graphics:      graphics,
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

func (game *Game) SetGraphicsLib(graphics flat_game.IGraphics) {
	game.graphics = graphics
}

func (game *Game) Graphics() flat_game.IGraphics {
	return game.graphics
}

func (game *Game) Config() flat_game.Config {
	return game.config
}

func (game *Game) preTick() {
	game.graphics.PreTick()
}

func (game *Game) Tick(delta float32) {
	game.graphics.Tick()

	runTick(game, nil, game.CurrentScene(), delta)
}

func (game *Game) postTick() {
	game.graphics.PostTick()
}

func (game *Game) onKeyEvent(key input.Key, event input.KeyEvent) {
	game.CurrentScene().OnKeyEvent(key, event)
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
