package flat_game

type IExt interface {
	CanTick(game IGame) bool

	Tick(game IGame, delta float32)
}
