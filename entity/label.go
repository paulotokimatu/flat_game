package entity

import (
	"github.com/paulotokimatu/flat_game"

	"github.com/paulotokimatu/flat_game/utils"
)

type LabelEnt struct {
	flat_game.IEntity
	color *utils.Vec3
	font  flat_game.IFont
	text  string
}

func NewLabelEnt(config *Config, font flat_game.IFont, text string, color *utils.Vec3) *LabelEnt {
	base := NewBaseEntity(config)

	return &LabelEnt{
		IEntity: base,
		color:   color,
		font:    font,
		text:    text,
	}
}

func (ent *LabelEnt) Tick(game flat_game.IGame, parent flat_game.IEntity, delta float32) {
	game.Graphics().DrawLabel(
		ent.font,
		ent.text,
		ent.Position(),
		ent.color,
	)
}

func (ent *LabelEnt) SetText(text string) {
	ent.text = text
}
