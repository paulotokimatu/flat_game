package game_test

import (
	"flat_game/entity"
	"flat_game/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetSceneShouldAddSceneIfNotPresent(t *testing.T) {
	scenesManager := game.NewScenesManager()

	assert.Nil(t, scenesManager.SceneByName("foo"))

	scenesManager.SetScene(entity.NewScene(&entity.Config{Name: "foo"}), false)

	assert.NotNil(t, scenesManager.SceneByName("foo"))
	assert.Equal(t, "foo", scenesManager.CurrentScene().Name())
}

func TestSetSceneShouldRemovePreviousSceneIfFlagged(t *testing.T) {
	scenesManager := game.NewScenesManager()

	fooScene := entity.NewScene(&entity.Config{Name: "foo"})

	scenesManager.AddScene(fooScene)

	assert.NotNil(t, scenesManager.SceneByName("foo"))
	scenesManager.SetScene(fooScene, false)

	scenesManager.SetScene(entity.NewScene(&entity.Config{Name: "bar"}), true)

	assert.Nil(t, scenesManager.SceneByName("foo"))
	assert.Equal(t, "bar", scenesManager.CurrentScene().Name())
}
