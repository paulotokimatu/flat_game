package graphics

import (
	"flat_game"
	"fmt"
	"image"
	"image/draw"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Glyph struct {
	advance  int
	bearingH int
	bearingV int
	height   int
	texture  flat_game.ITexture
	width    int
}

func NewGlyph(truetypeFont *truetype.Font, charCode rune, scale int) (*Glyph, error) {
	truetypeFace := truetype.NewFace(truetypeFont, &truetype.Options{
		Size:    float64(scale),
		DPI:     72,
		Hinting: font.HintingFull,
	})

	gBnd, gAdv, ok := truetypeFace.GlyphBounds(charCode)
	if !ok {
		return nil, fmt.Errorf("ttf face glyphBounds error")
	}

	gh := int32((gBnd.Max.Y - gBnd.Min.Y) >> 6)
	gw := int32((gBnd.Max.X - gBnd.Min.X) >> 6)

	//if glyph has no dimensions set to a max value
	if gw == 0 || gh == 0 {
		gBnd = truetypeFont.Bounds(fixed.Int26_6(scale))
		gw = int32((gBnd.Max.X - gBnd.Min.X) >> 6)
		gh = int32((gBnd.Max.Y - gBnd.Min.Y) >> 6)

		//above can sometimes yield 0 for font smaller than 48pt, 1 is minimum
		if gw == 0 || gh == 0 {
			gw = 1
			gh = 1
		}
	}

	//The glyph's ascent and descent equal -bounds.Min.Y and +bounds.Max.Y.
	gAscent := int(-gBnd.Min.Y) >> 6
	gdescent := int(gBnd.Max.Y) >> 6

	//create image to draw glyph
	fg := image.White
	bg := image.Black
	rect := image.Rect(0, 0, int(gw), int(gh))
	rgba := image.NewRGBA(rect)
	draw.Draw(rgba, rgba.Bounds(), bg, image.Point{}, draw.Src)

	freetypeContext := createFreetypeContext(truetypeFont, scale, rgba, fg)

	//set the glyph dot
	px := 0 - (int(gBnd.Min.X) >> 6)
	py := (gAscent)
	pt := freetype.Pt(px, py)

	// Draw the text from mask to image
	_, err := freetypeContext.DrawString(string(charCode), pt)
	if err != nil {
		return nil, err
	}

	texture := NewTexture("glyph_"+string(charCode), int32(rgba.Rect.Dx()), int32(rgba.Rect.Dy()), rgba.Pix)
	// I dont know if above work since the WRAP props are different:
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	// gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	glyph := &Glyph{
		advance:  int(gAdv),
		bearingH: (int(gBnd.Min.X) >> 6),
		bearingV: gdescent,
		height:   int(gh),
		texture:  texture,
		width:    int(gw),
	}

	return glyph, nil
}

func createFreetypeContext(truetypeFont *truetype.Font, scale int, rgba *image.RGBA, fg *image.Uniform) *freetype.Context {
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(truetypeFont)
	c.SetFontSize(float64(scale))
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)
	c.SetHinting(font.HintingFull)

	return c
}

func (glyph *Glyph) Vertices(drawPositionX float32, drawPositionY float32) []float32 {
	xPos := drawPositionX + float32(glyph.bearingH)
	yPos := drawPositionY - float32(glyph.height-glyph.bearingV)
	width := float32(glyph.width)
	height := float32(glyph.height)

	return []float32{
		xPos + width, yPos, 1.0, 0.0,
		xPos, yPos, 0.0, 0.0,
		xPos, yPos + height, 0.0, 1.0,

		xPos, yPos + height, 0.0, 1.0,
		xPos + width, yPos + height, 1.0, 1.0,
		xPos + width, yPos, 1.0, 0.0,
	}
}

func (glyph *Glyph) Advance() int {
	return glyph.advance
}

func (glyph *Glyph) Texture() flat_game.ITexture {
	return glyph.texture
}
