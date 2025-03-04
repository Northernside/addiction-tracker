package commands

import (
	"fmt"
)

var ErrNotEnoughArgs = fmt.Errorf("Not enough arguments")
var StreakGoalNotANumber = fmt.Errorf("Streak goal is not an integer")

func Help(args []string) error {
	for _, cmd := range Commands {
		fmt.Printf("Command: %s", cmd.Keys[0])
		fmt.Printf("\nDescription: %s\n\n", cmd.Desc)
	}

	return nil
}
