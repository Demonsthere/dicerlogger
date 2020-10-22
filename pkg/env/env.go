package env

import (
	"github.com/vrischmann/envconfig"
)

var (
	Config EnvConfig
)

//EnvConfig .
type EnvConfig struct {
	DiceSize     int `envconfig:"default=100"`
	TimeInterval int `envconfig:"default=1000"`
}

//InitConfig initialize config from envs
func InitConfig() {
	err := envconfig.Init(&Config)
	if err != nil {
		panic(err)
	}
}
