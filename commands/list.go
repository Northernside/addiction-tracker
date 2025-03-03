package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func List(args []string) error {
	// usage: list

	// load all addictions
	addictions, err := LoadAddictions()
	if err != nil {
		return err
	}

	if len(addictions) == 0 {
		fmt.Println("Keine Sucht vorhanden")
	}

	for _, addiction := range addictions {
		// calculate the streak correctly
		streak := int(time.Since(addiction.StartedAt).Hours() / 24)

		fmt.Printf("Sucht:\t\t%s\n", addiction.Name)
		fmt.Printf("Gestartet am:\t%s um %s\n\n", addiction.StartedAt.Format("02.01.2006"), addiction.StartedAt.Format("15:04"))

		fmt.Printf("Streak:\t\t%d 🔥\n", streak)
		fmt.Printf("Ziel:\t\t%d 🔥\n\n", addiction.StreakGoal)

		// if streak goal is reached, print message to ask the user if they want to increase the streak goal or not
		if streak >= addiction.StreakGoal {
			fmt.Printf("Streak Ziel erreicht! %d/%d Tage. Möchtest du den Streak erhöhen?\n", streak, addiction.StreakGoal)

			// get input from user
			var response string
			fmt.Print("Antwort (ja/nein): ")
			fmt.Scanln(&response)
			if response == "ja" {
				// ask user for new streak goal
				var newStreakGoalStr string

				fmt.Print("Neues Streak Ziel: ")
				fmt.Scanln(&newStreakGoalStr)

				newStreakGoal, err := strconv.Atoi(newStreakGoalStr)
				if err != nil {
					return StreakGoalNotANumber
				}

				// update addiction
				addiction.StreakGoal = newStreakGoal
				UpdateAddiction(addiction)
			}
		}

		// print a progress bar
		progress(streak, addiction.StreakGoal)
	}

	return nil
}

func progress(streak int, streakGoal int) {
	const barWidth = 32

	fullBlocks := (streak * barWidth) / streakGoal

	// fractional / remainder part of the progress
	fractionalPart := (streak * barWidth) % streakGoal
	partialBlock := ""
	if fractionalPart > 0 {
		partialBlock = getPartialBlock(fractionalPart, streakGoal)
	}

	emptyBlocks := barWidth - fullBlocks - len(partialBlock)

	if emptyBlocks < 0 {
		emptyBlocks = 0
	}

	var daysStr = "Tag"
	if streak != 1 {
		daysStr += "e"
	}

	fmt.Printf("[%s%s%s] %d/%d %s\n",
		strings.Repeat("█", fullBlocks),
		partialBlock,
		strings.Repeat("░", emptyBlocks),
		streak, streakGoal, daysStr)
}

func getPartialBlock(fractionalPart int, streakGoal int) string {
	fraction := float64(fractionalPart) / float64(streakGoal)
	switch {
	case fraction >= 7/8:
		return "▉"
	case fraction >= 3/4:
		return "▊"
	case fraction >= 5/8:
		return "▋"
	case fraction >= 1/2:
		return "▌"
	case fraction >= 3/8:
		return "▍"
	case fraction >= 1/4:
		return "▎"
	case fraction >= 1/8:
		return "▏"
	default:
		return ""
	}
}
