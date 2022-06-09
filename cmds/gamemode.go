package cmds

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type GamemodeCommandSpec struct {
	Gamemode gamemodeSpec `name:"gameMode"`
	Target   []cmd.Target `name:"player" optional:""`
}

type GamemodeCommandInt struct {
	Gamemode int          `name:"gameMode"`
	Target   []cmd.Target `name:"player" optional:""`
}

func (cmd GamemodeCommandSpec) Run(source cmd.Source, output *cmd.Output) {
	var target *player.Player
	if len(cmd.Target) < 1 {
		target = source.(*player.Player)
	} else {
		t, ok := cmd.Target[0].(*player.Player)
		if !ok {
			output.Error("Selector must be player-type")
		}
		target = t
	}
	if cmd.Gamemode == "d" || cmd.Gamemode == "default" {
		target.SetGameMode(target.World().DefaultGameMode())
	} else {
		target.SetGameMode(stringToGamemode(string(cmd.Gamemode)))
	}
}

func (cmd GamemodeCommandInt) Run(source cmd.Source, output *cmd.Output) {}

type gamemodeSpec string

func (gamemodeSpec) Type() string {
	return "GameMode"
}

func (gamemodeSpec) Options(source cmd.Source) []string {
	return []string{
		"a", "adventure",
		"c", "creative",
		"d", "default",
		"s", "survival",
		"spectator",
	}
}

func stringToGamemode(input string) world.GameMode {
	var gamemode world.GameMode
	switch input {
	case "a", "adventure":
		gamemode = world.GameModeAdventure
	case "c", "creative":
		gamemode = world.GameModeCreative
	case "s", "survival":
		gamemode = world.GameModeSurvival
	case "spectator":
		gamemode = world.GameModeSpectator
	default:
		gamemode = world.GameModeCreative
	}
	return gamemode
}
