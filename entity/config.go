package entity

import (
	"github.com/paulotokimatu/flat_game"

	"github.com/paulotokimatu/flat_game/utils"
)

type Config struct {
	Name     string
	Position utils.Vec2
	Size     utils.Vec2
	Children []flat_game.IEntity
}
