package commands

import (
	"fmt"
)

var ErrNotEnoughArgs = fmt.Errorf("Nicht genügend Argumente")
var StreakGoalNotANumber = fmt.Errorf("Streak Goal ist kein Integer")

func Help(args []string) error {
	for _, cmd := range Commands {
		fmt.Printf("Befehl: %s", cmd.Keys[0])
		fmt.Printf("\nBeschreibung: %s\n\n", cmd.Desc)
	}

	return nil
}
