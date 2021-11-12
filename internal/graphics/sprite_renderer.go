package graphics

import (
	"github.com/paulotokimatu/flat_game"

	"github.com/paulotokimatu/flat_game/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type SpriteRenderer struct {
	shader *Shader
	Vao    uint32
}

var vertices = []float32{
	// pos    // tex
	0.0, 1.0, 0.0, 1.0,
	1.0, 0.0, 1.0, 0.0,
	0.0, 0.0, 0.0, 0.0,

	0.0, 1.0, 0.0, 1.0,
	1.0, 1.0, 1.0, 1.0,
	1.0, 0.0, 1.0, 0.0,
}

func NewSpriteRenderer(shader *Shader) (*SpriteRenderer, error) {
	var vao uint32

	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	vertexAttrib := uint32(gl.GetAttribLocation(shader.ProgramId, gl.Str("vertex\x00")))
	gl.EnableVertexAttribArray(vertexAttrib)
	gl.VertexAttribPointerWithOffset(vertexAttrib, 4, gl.FLOAT, false, 4*4, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return &SpriteRenderer{shader, vao}, nil
}

func (renderer *SpriteRenderer) DrawSprite(texture flat_game.ITexture, position *utils.Vec2, size *utils.Vec2, color *utils.Vec3) {
	// model = mgl32.Translate2D(model, mgl32.Vec3(position, 0.0))
	// model = mgl32.Translate2D(model, mgl32.Vec3(0.5 * size.x, 0.5 * size.y, 0.0))

	renderer.shader.Use()

	model := mgl32.Ident4()
	model = model.Mul4(mgl32.Translate3D(position.X, position.Y, 1))
	model = model.Mul4(mgl32.Scale3D(size.X, size.Y, 1))
	renderer.shader.SetMatrix4("model", &model, false)

	colorMlgVector := mgl32.Vec3{color.X, color.Y, color.Z}

	renderer.shader.SetVector3f("spriteColor", colorMlgVector, false)

	gl.ActiveTexture(gl.TEXTURE0)

	texture.Bind()

	gl.BindVertexArray(renderer.Vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
	gl.BindVertexArray(0)
}

// func (renderer *SpriteRenderer) cleanUp() {
// 	gl.DeleteVertexArrays(1, &renderer.vao)
// }
