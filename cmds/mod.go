package cmds

import (
	"github.com/RoyalMCPE/df-vanillacmds/cmds/time"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
)

var srv *server.Server

func RegisterVanillaCommands(server *server.Server) {
	srv = server
	cmd.Register(cmd.New("time", "Changes or queries the world's game time.", []string{}, time.TimeSetString{}, time.TimeSetInt{}, time.TimeAdd{}))
	cmd.Register(cmd.New("weather", "Sets the weather.", []string{}, WeatherCommand{}))
}
