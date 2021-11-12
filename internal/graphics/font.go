package graphics

import (
	"github.com/paulotokimatu/flat_game"

	"github.com/paulotokimatu/flat_game/utils"

	"github.com/golang/freetype/truetype"
)

type Font struct {
	charCodeMin rune
	glyphs      []flat_game.IGlyph
}

func NewFont(name string, fontData []byte, charCodeMin rune, charCodeMax rune, scale int) (*Font, error) {
	trueTypeFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	glyphs := make([]flat_game.IGlyph, 0, charCodeMax-charCodeMin+1)

	for charCord := charCodeMin; charCord <= charCodeMax; charCord++ {
		glyph, err := NewGlyph(trueTypeFont, charCord, scale)
		if err != nil {
			return nil, err
		}

		glyphs = append(glyphs, glyph)
	}

	font := &Font{
		charCodeMin: charCodeMin,
		glyphs:      glyphs,
	}

	return font, nil
}

func NewFontFromFile(name string, fileName string, charCodeMin rune, charCodeMax rune, scale int) (*Font, error) {
	fontData, err := utils.ReadTextFile(fileName)
	if err != nil {
		return nil, err
	}

	return NewFont(name, fontData, charCodeMin, charCodeMax, scale)
}

func (font *Font) Glyph(runeIndex rune) flat_game.IGlyph {
	// skip runes that are not in font char range
	if int(runeIndex)-int(font.charCodeMin) > len(font.glyphs) || runeIndex < font.charCodeMin {
		return nil
	}

	glyph := font.glyphs[runeIndex-font.charCodeMin]

	return glyph
}
