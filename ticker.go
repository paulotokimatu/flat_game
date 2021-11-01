package flat_game

type ITicker interface {
	CanTick(game IGame) bool

	Tick(game IGame, delta float32)
}
