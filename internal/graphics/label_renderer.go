package graphics

import (
	"github.com/paulotokimatu/flat_game"

	"github.com/paulotokimatu/flat_game/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type LabelRenderer struct {
	shader *Shader
	vao    uint32
	vbo    uint32
}

func NewLabelRenderer(shader *Shader) (*LabelRenderer, error) {
	shader.Use()

	//set screen resolution
	shader.SetVector2f("resolution", mgl32.Vec2{float32(800), float32(600)}, false)

	gl.BindTexture(gl.TEXTURE_2D, 0)

	var vao uint32
	var vbo uint32

	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.BindVertexArray(vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	gl.BufferData(gl.ARRAY_BUFFER, 6*4*4, nil, gl.STATIC_DRAW)

	vertexAttrib := uint32(gl.GetAttribLocation(shader.ProgramId, gl.Str("vertex\x00")))
	gl.EnableVertexAttribArray(vertexAttrib)
	gl.VertexAttribPointer(vertexAttrib, 4, gl.FLOAT, false, 4*4, gl.PtrOffset(0))
	defer gl.DisableVertexAttribArray(vertexAttrib)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	renderer := &LabelRenderer{
		shader: shader,
		vao:    vao,
		vbo:    vbo,
	}

	return renderer, nil
}

func (renderer *LabelRenderer) Draw(
	font flat_game.IFont,
	text string,
	startDrawX float32,
	startDrawY float32,
	color *utils.Vec3,
) {
	if len(text) == 0 {
		return
	}

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	renderer.shader.Use()

	colorMlgVector := mgl32.Vec4{color.X, color.Y, color.Z, 1.0}
	renderer.shader.SetVector4f("labelColor", colorMlgVector, false)

	gl.BindVertexArray(renderer.vao)
	gl.ActiveTexture(gl.TEXTURE0)

	drawX := startDrawX
	for _, runeIndex := range text {
		glyph := font.Glyph(runeIndex)

		if glyph == nil {
			continue
		}

		vertices := glyph.Vertices(drawX, startDrawY)

		renderer.drawRune(glyph.Texture(), vertices)

		// Bitshift by 6 to get value in pixels (2^6 = 64 (divide amount of 1/64th pixels by 64 to get amount of pixels))
		drawX += float32((glyph.Advance() >> 6))
	}

	gl.BindVertexArray(0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.UseProgram(0)
	gl.Disable(gl.BLEND)
}

func (renderer *LabelRenderer) drawRune(texture flat_game.ITexture, vertices []float32) {
	texture.Bind()

	gl.BindBuffer(gl.ARRAY_BUFFER, renderer.vbo)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(vertices)*4, gl.Ptr(vertices))
	gl.DrawArrays(gl.TRIANGLES, 0, 16)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
