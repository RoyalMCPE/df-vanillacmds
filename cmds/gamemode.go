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

func (c GamemodeCommandSpec) Run(source cmd.Source, output *cmd.Output) {
	var selector, ok = c.Target.LoadOr([]cmd.Target{source})[0].(*player.Player)
	if !ok {
		output.Error("Selector must be player-type")
		return
	}
	if c.Gamemode == "d" || c.Gamemode == "default" {
		selector.SetGameMode(selector.World().DefaultGameMode())
		if selector.Name() != source.Name() {
			output.Printf("Set %v's game mode to Default", selector.Name())
			selector.Messagef("Your game mode has been updated to Default")
			return
		}
		output.Printf("Your game mode has been updated to Default")
		output.Printf("Set own game mode to Default")
	} else {
		var gamemode, name = stringToGamemode(string(c.Gamemode))
		selector.SetGameMode(gamemode)
		if selector.Name() != source.Name() {
			output.Printf("Set %v's game mode to %v", selector.Name(), name)
			selector.Messagef("Your game mode has been updated to %v", name)
			return
		}
		output.Printf("Your game mode has been updated to %v", name)
		output.Printf("Set own game mode to %v", name)
	}
}

func (c GamemodeCommandInt) Run(source cmd.Source, output *cmd.Output) {
	var selector, ok = c.Target.LoadOr([]cmd.Target{source})[0].(*player.Player)
	if !ok {
		output.Error("Selector must be player-type")
		return
	}

	if c.Gamemode == 5 {
		selector.SetGameMode(selector.World().DefaultGameMode())
		if selector.Name() != source.Name() {
			output.Printf("Set %v's game mode to Default", selector.Name())
			selector.Message("Your game mode has been updated to Default")
			return
		}
		output.Print("Your game mode has been updated to Default")
		output.Print("Set own game mode to Default")
	} else if gamemode, name := intToGamemode(c.Gamemode); gamemode != nil {
		selector.SetGameMode(gamemode)
		if selector.Name() != source.Name() {
			output.Printf("Set %v's game mode to %v", selector.Name(), name)
			selector.Messagef("Your game mode has been update to %v", name)
			return
		}
		output.Printf("Your game mode has been update to %v", name)
		output.Printf("Set own game mode to %v", name)
	} else {
		output.Errorf("Game mode '%v' is invalid", c.Gamemode)
	}
}

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
}

func intToGamemode(input int) (world.GameMode, string) {
	switch input {
	case 0:
		return world.GameModeSurvival, "Survival"
	case 1:
		return world.GameModeCreative, "Creative"
	case 2:
		return world.GameModeAdventure, "Adventure"
	case 6:
		return world.GameModeSpectator, "Spectator"
	default:
		return nil, ""
	}
}
