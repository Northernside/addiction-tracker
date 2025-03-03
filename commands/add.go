package commands

import (
	"fmt"
	"strconv"
	"time"
)

func Add(args []string) error {
	// usage: add <addiction name> <streak goal>

	if len(args) < 2 {
		fmt.Println("Benutzung: add <addiction name> <streak goal>")
		return ErrNotEnoughArgs
	}

	name := args[0]
	streakGoalStr := args[1]

	// check if streak goal is a number
	streakGoal, err := strconv.Atoi(streakGoalStr)
	if err != nil {
		return StreakGoalNotANumber
	}

	addiction := Addiction{
		Name:       name,
		StreakGoal: streakGoal,
		StartedAt:  time.Now().Add(-48 * time.Hour),
	}

	// save addiction to file
	err = SaveAddiction(addiction)
	if err != nil {
		return err
	}

	fmt.Printf("Sucht %s wurde hinzugefügt!\n", name)

	return nil
}
