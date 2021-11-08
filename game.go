package flat_game

type IGame interface {
	IAssetsManager

	ISceneManager

	AddFont(name string, fileName string, charCodeMin rune, charCodeMax rune, scale int) (IFont, error)

	AddTexture(name string, fileName string) (ITexture, error)

	Config() Config

	Graphics() IGraphics

	TextureByName(name string) ITexture
}

type ISceneManager interface {
	AddScene(scene IScene)

	CurrentScene() IScene

	DeleteScene(sceneName string)

	SceneByName(sceneName string) IScene

	SetScene(scene IScene, deletePreviousScene bool)
}

type IAssetsManager interface {
	AddFont(name string, fileName string, charCodeMin rune, charCodeMax rune, scale int) (IFont, error)

	AddTexture(name string, fileName string) (ITexture, error)

	FontByName(name string) IFont

	TextureByName(name string) ITexture
}
