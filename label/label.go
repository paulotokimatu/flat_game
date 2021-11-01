package label

import (
	"flat_game"
	"flat_game/utils"
)

type SpriteExt struct {
	parent  flat_game.IEntity // instead of saving the parent, better pass as argument?
	texture flat_game.ITexture
}

func NewLabel(parent flat_game.IEntity, texture flat_game.ITexture) *SpriteExt {
	return &SpriteExt{
		parent:  parent,
		texture: texture,
	}
}

func (ext *SpriteExt) CanTick(game flat_game.IGame) bool {
	return true
}

func (ext *SpriteExt) Tick(game flat_game.IGame, delta float32) {
	game.Graphics().DrawSprite(
		ext.texture,
		ext.parent.Position(),
		ext.parent.Size(),
		utils.Vec3{
			X: 1.0,
			Y: 1.0,
			Z: 1.0,
		},
	)
}
