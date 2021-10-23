package graphics

import (
	"flat_game/utils"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Texture struct {
	Name string
	Id   uint32
}

func NewTexture(name string, width int32, height int32, pixels []uint8) *Texture {
	var textureId uint32

	gl.GenTextures(1, &textureId)
	gl.ActiveTexture(gl.TEXTURE0)

	gl.BindTexture(gl.TEXTURE_2D, textureId)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, width, height, 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(pixels))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	// gl.BindTexture(gl.TEXTURE_2D, 0)

	return &Texture{name, textureId}
}

func NewTextureFromFile(name string, fileName string) *Texture {
	rgba, err := utils.ReadImage(fileName)
	if err != nil {
		panic(err)
	}

	imageWidth := int32(rgba.Rect.Size().X)
	imageHeight := int32(rgba.Rect.Size().Y)

	return NewTexture(name, imageWidth, imageHeight, rgba.Pix)
}

func (texture *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, texture.Id)
}
