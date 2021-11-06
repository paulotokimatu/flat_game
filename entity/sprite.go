package entity

import (
	"flat_game"
	"flat_game/utils"
)

type SpriteEnt struct {
	flat_game.IEntity
	texture               flat_game.ITexture
	useParentSizePosition bool
}

func NewSpriteEnt(config *Config, texture flat_game.ITexture, useParentSizePosition bool) *SpriteEnt {
	base := NewBaseEntity(config)

	return &SpriteEnt{base, texture, useParentSizePosition}
}

func (ent *SpriteEnt) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	var position *utils.Vec2
	var size *utils.Vec2

	if ent.useParentSizePosition {
		position = parent.Position()
		size = parent.Size()
	} else {
		position = ent.Position()
		size = ent.Size()
	}

	game.Graphics().DrawSprite(
		ent.texture,
		position,
		size,
		&utils.Vec3{
			X: 1.0,
			Y: 1.0,
			Z: 1.0,
		},
	)
}
