package cmds

import (
	"github.com/RoyalMCPE/df-vanillacmds/cmds/time"
	"github.com/df-mc/dragonfly/server/cmd"
)

func RegisterVanillaCommands() {
	cmd.Register(cmd.New("time", "Changes or queries the world's game time.", []string{}, time.TimeSetString{}, time.TimeSetInt{}, time.TimeAdd{}))
	cmd.Register(cmd.New("weather", "Sets the weather.", []string{}, WeatherCommand{}))
}
