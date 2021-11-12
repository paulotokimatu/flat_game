package game

import (
	"github.com/paulotokimatu/flat_game"
)

type ScenesManager struct {
	currentScene flat_game.IScene
	scenes       map[string]flat_game.IScene
}

func NewScenesManager() *ScenesManager {
	return &ScenesManager{
		scenes: map[string]flat_game.IScene{},
	}
}

func (game *ScenesManager) AddScene(scene flat_game.IScene) {
	if game.scenes == nil {
		game.scenes = map[string]flat_game.IScene{}
	}

	game.scenes[scene.Name()] = scene
}

func (game *ScenesManager) CurrentScene() flat_game.IScene {
	return game.currentScene
}

func (game *ScenesManager) DeleteScene(sceneName string) {
	if game.CurrentScene().Name() == sceneName {
		panic("cannot delete current scene")
	}

	delete(game.scenes, sceneName)
}

func (game *ScenesManager) SceneByName(sceneName string) flat_game.IScene {
	return game.scenes[sceneName]
}

func (game *ScenesManager) SetScene(scene flat_game.IScene, deletePreviousScene bool) {
	previousScene := game.CurrentScene()

	if _, ok := game.scenes[scene.Name()]; !ok {
		game.AddScene(scene)
	}

	game.currentScene = scene

	if previousScene != nil && deletePreviousScene {
		game.DeleteScene(previousScene.Name())
	}
}
