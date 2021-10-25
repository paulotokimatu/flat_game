package game

import (
	"flat_game"
	"flat_game/internal/graphics"
	"fmt"
)

type GraphicsLib string

const (
	OpenGlLib GraphicsLib = "opengl"
)

func NewGraphics(lib GraphicsLib) (flat_game.IGraphics, error) {
	if lib == OpenGlLib {
		return &graphics.OpenGl{}, nil
	}

	return nil, fmt.Errorf("invalid graphic library %q", lib)
}
