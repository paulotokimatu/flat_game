package flat_game

type IGame interface {
	AddFont(name string, fileName string, charCodeMin rune, charCodeMax rune, scale int) (IFont, error)

	AddTexture(name string, fileName string) (ITexture, error)

	Config() Config

	Graphics() IGraphics

	Scene() IScene

	SetScene(scene IScene)

	TextureByName(name string) ITexture
}
