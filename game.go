package flat_game

type IGame interface {
	AddTexture(name string, fileName string) (ITexture, error)

	Config() Config

	Graphics() IGraphics

	Scene() IScene

	SetScene(scene IScene)

	TextureByName(name string) ITexture
}
