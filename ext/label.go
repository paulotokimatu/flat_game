package ext

import (
	"flat_game"
	"flat_game/utils"
)

type LabelExt struct {
	parent flat_game.IEntity
	font   flat_game.IFont
	text   string
}

func NewLabelExt(parent flat_game.IEntity, font flat_game.IFont, text string) *LabelExt {
	return &LabelExt{
		parent: parent,
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
		utils.Vec3{
			X: 1.0,
			Y: 1.0,
			Z: 1.0,
		},
	)
}

func (ext *LabelExt) SetText(text string) {
	ext.text = text
}
