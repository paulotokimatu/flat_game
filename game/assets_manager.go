package game

import (
	"flat_game"
	graphicsLib "flat_game/internal/graphics"
)

type AssetsManager struct {
	fonts    map[string]flat_game.IFont
	textures map[string]flat_game.ITexture
}

func NewAssetsManager() *AssetsManager {
	return &AssetsManager{
		fonts:    map[string]flat_game.IFont{},
		textures: map[string]flat_game.ITexture{},
	}
}

func (game *AssetsManager) AddTexture(name string, fileName string) (flat_game.ITexture, error) {
	texture, err := graphicsLib.NewTextureFromFile(name, fileName)
	if err != nil {
		return texture, err
	}

	game.textures[texture.Name] = texture

	return texture, nil
}

func (game *AssetsManager) TextureByName(name string) flat_game.ITexture {
	return game.textures[name]
}

func (game *AssetsManager) AddFont(
	name string,
	fileName string,
	charCodeMin rune,
	charCodeMax rune,
	scale int,
) (flat_game.IFont, error) {
	font, err := graphicsLib.NewFontFromFile(name, fileName, charCodeMin, charCodeMax, scale)
	if err != nil {
		return nil, err
	}

	game.fonts[name] = font

	return font, nil
}

func (game *AssetsManager) FontByName(name string) flat_game.IFont {
	return game.fonts[name]
}
