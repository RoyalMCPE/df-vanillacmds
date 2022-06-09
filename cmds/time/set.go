package time

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

type TimeSetInt struct {
	Sub set_name

	Time int `cmd:"amount"`
}

type TimeSetString struct {
	Sub set_name

	Time time_spec `cmd:"time"`
}

func (cmd TimeSetString) Run(source cmd.Source, output *cmd.Output) {
	time := timeFromString(string(cmd.Time))
	source.World().SetTime(time)
	output.Printf("Set the time to %v", time)
}

func (cmd TimeSetInt) Run(source cmd.Source, output *cmd.Output) {
	source.World().SetTime(cmd.Time)
	output.Printf("Set the time to %v", cmd.Time)
}

func timeFromString(time string) int {
	return map[string]int{
		"day": 1000, "night": 13000, "noon": 6000, "midnight": 18000, "sunrise": 23000, "sunset": 12000,
	}[time]
}

type set_name string

func (set_name) SubName() string {
	return "set"
}

type time_spec string

func (time_spec) Type() string {
	return "TimeSpec"
}

func (time_spec) Options(source cmd.Source) []string {
	return []string{"day", "night", "noon", "midnight", "sunrise", "sunset"}
}
