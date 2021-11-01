package ext

import (
	"flat_game"
	"flat_game/utils"
)

type LabelExt struct {
	color  *utils.Vec3
	font   flat_game.IFont
	parent flat_game.IEntity
	text   string
}

func NewLabelExt(parent flat_game.IEntity, font flat_game.IFont, text string, color *utils.Vec3) *LabelExt {
	return &LabelExt{
		parent: parent,
		color:  color,
		font:   font,
		text:   text,
	}
}

func (ext *LabelExt) CanTick(game flat_game.IGame) bool {
	return true
}

func (ext *LabelExt) Tick(game flat_game.IGame, delta float32) {
	game.Graphics().DrawLabel(
		ext.font,
		ext.text,
		ext.parent.Position(),
		ext.color,
	)
}

func (ext *LabelExt) SetText(text string) {
	ext.text = text
}
