package time

import "github.com/df-mc/dragonfly/server/cmd"

type TimeAdd struct {
	Sub add_name

	Amount int `cmd:"amount"`
}

func (cmd TimeAdd) Run(source cmd.Source, output *cmd.Output) {
	source.World().SetTime(source.World().Time() + cmd.Amount)
	output.Printf("Added 10 to the time")
}

type add_name string

func (add_name) SubName() string {
	return "add"
}
