package time

import "github.com/df-mc/dragonfly/server/cmd"

type TimeAdd struct {
	Sub cmd.SubCommand `cmd:"add"`

	Amount int `cmd:"amount"`
}

func (cmd TimeAdd) Run(source cmd.Source, output *cmd.Output) {
	source.World().SetTime(source.World().Time() + cmd.Amount)
	output.Printf("Added 10 to the time")
}
