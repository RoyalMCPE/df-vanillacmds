package cmds

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"math/rand"
	"time"
)

type WeatherCommand struct {
	State    weather `cmd:"state"`
	Duration cmd.Optional[int]
}

func (cmd WeatherCommand) Run(source cmd.Source, output *cmd.Output) {
	var dur = cmd.Duration.LoadOr(0)
	var duration time.Duration = time.Duration(dur) * time.Second
	if dur == 0 {
		duration = time.Duration(rand.Intn(900-300)+300) * time.Second
	}
	if cmd.State == "thunder" {
		source.World().StartThundering(duration)
		output.Print("Changing to rain and thunder")
	} else if cmd.State == "rain" {
		source.World().StartRaining(duration)
		output.Print("Changing to rainy weather")
	} else {
		source.World().StopThundering()
		source.World().StopRaining()
		output.Print("Changing to clear weather")
	}
}

type weather string

func (weather) Type() string {
	return "WeatherState"
}

func (weather) Options(source cmd.Source) []string {
	return []string{"clear", "rain", "thunder"}
}
