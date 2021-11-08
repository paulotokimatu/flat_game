package flat_game

type IGame interface {
	AddFont(name string, fileName string, charCodeMin rune, charCodeMax rune, scale int) (IFont, error)

	AddTexture(name string, fileName string) (ITexture, error)

	AddScene(scene IScene)

	Config() Config

	CurrentScene() IScene

	Graphics() IGraphics

	DeleteScene(sceneName string)

	SceneByName(sceneName string) IScene

	SetScene(scene IScene, deletePreviousScene bool)

	TextureByName(name string) ITexture
}
