package flat_game

import (
	"encoding/json"

	"github.com/paulotokimatu/flat_game/utils"
)

type Config struct {
	BgColor  utils.Vec3 `json:"bg_color"`
	Graphics string     `json:"graphics"`
	MaxFPS   int        `json:"max_fps"`
	Name     string     `json:"name"`
	Size     utils.Vec2 `json:"size"`
}

func NewConfigFromJson(fileName string) Config {
	var config Config

	rawConfig, err := utils.ReadJsonFile(fileName)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(rawConfig, &config)

	return config
}
