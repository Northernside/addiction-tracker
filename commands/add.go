package commands

import (
	"fmt"
	"strconv"
	"time"
)

func Add(args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: add <addiction name> (streak goal)")
		return ErrNotEnoughArgs
	}

	name := args[0]

	var streakGoal int = -1
	var err error

	if len(args) > 1 {
		streakGoalStr := args[1]

		// check if streak goal is a number
		streakGoal, err = strconv.Atoi(streakGoalStr)
		if err != nil {
			return StreakGoalNotANumber
		}
	}

	addiction := Addiction{
		Name:       name,
		StreakGoal: streakGoal,
		StartedAt:  time.Now(),
	}

	// save addiction to file
	err = SaveAddiction(addiction)
	if err != nil {
		return err
	}

	fmt.Printf("Added '%s' addiction", name)
	if streakGoal != -1 {
		fmt.Printf("with a streak goal of %d", streakGoal)
	}

	fmt.Println()

	return nil
}
