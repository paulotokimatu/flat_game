package flat_game

import "flat_game/utils"

type GameConfig struct {
	BgColor  utils.Vec3 `json:"bg_color"`
	Graphics string     `json:"graphics"`
	MaxFPS   int        `json:"max_fps"`
	Name     string     `json:"name"`
	Size     utils.Vec2 `json:"size"`
}
