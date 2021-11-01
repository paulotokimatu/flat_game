package entity

import (
	"flat_game"
	"flat_game/utils"
)

type SpriteEnt struct {
	flat_game.IEntity
	texture flat_game.ITexture
}

func NewSpriteEnt(config *Config, texture flat_game.ITexture) *SpriteEnt {
	base := NewBaseEntity(config)

	return &SpriteEnt{base, texture}
}

func (ent *SpriteEnt) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	game.Graphics().DrawSprite(
		ent.texture,
		parent.Position(),
		parent.Size(),
		&utils.Vec3{
			X: 1.0,
			Y: 1.0,
			Z: 1.0,
		},
	)
}
