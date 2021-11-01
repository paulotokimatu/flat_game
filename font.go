package flat_game

type IGlyph interface {
	Advance() int

	Texture() ITexture

	Vertices(drawPositionX float32, drawPositionY float32) []float32
}

type IFont interface {
	Glyph(runeIndex rune) IGlyph
}
