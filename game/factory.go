package game

import (
	"flat_game"
	"flat_game/internal/graphics"
)

type GraphicsLib string

const (
	OpenGlLib GraphicsLib = "opengl"
)

func NewGraphics(lib GraphicsLib) flat_game.IGraphics {
	if lib == OpenGlLib {
		return &graphics.OpenGl{}
	}

	panic("Invalid graphic library " + lib)
}
