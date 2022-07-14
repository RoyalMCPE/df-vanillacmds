package cmds

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type GamemodeCommandSpec struct {
	Gamemode gamemodeSpec               `cmd:"gameMode"`
	Target   cmd.Optional[[]cmd.Target] `cmd:"player"`
}

type GamemodeCommandInt struct {
	Gamemode int                        `cmd:"gameMode"`
	Target   cmd.Optional[[]cmd.Target] `cmd:"player"`
}

func (cmd GamemodeCommandSpec) Run(source cmd.Source, output *cmd.Output) {
	var target *player.Player
	var selector, _ = cmd.Target.Load()
	if len(selector) < 1 {
		target = source.(*player.Player)
	} else {
		t, ok := selector[0].(*player.Player)
		if !ok {
			output.Error("Selector must be player-type")
		}
		target = t
	}
	if cmd.Gamemode == "d" || cmd.Gamemode == "default" {
		target.SetGameMode(target.World().DefaultGameMode())
		target.Messagef("Your game mode has been updated to %v")
	} else {
		var gamemode, name = stringToGamemode(string(cmd.Gamemode))
		target.SetGameMode(gamemode)
		target.Messagef("Your game mode has been updated to %v", name)
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

func stringToGamemode(input string) (world.GameMode, string) {
	///var gamemode world.GameMode
	switch input {
	case "a", "adventure":
		return world.GameModeAdventure, "Adventure"
	case "c", "creative":
		return world.GameModeCreative, "Creative"
	case "s", "survival":
		return world.GameModeSurvival, "Survival"
	case "spectator":
		return world.GameModeSpectator, "Spectator"
	default:
		return world.GameModeCreative, "Creative"
	}
	//return gamemode
}
