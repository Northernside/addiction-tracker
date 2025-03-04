package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func List(args []string) error {
	addictions, err := LoadAddictions()
	if err != nil {
		return err
	}

	if len(addictions) == 0 {
		fmt.Println("No addictions found")
	}

	for _, addiction := range addictions {
		// calculate the streak correctly
		streak := int(time.Since(addiction.StartedAt).Hours() / 24)

		fmt.Printf("Addiction:\t%s\n", addiction.Name)
		fmt.Printf("Started on:\t%s at %s\n\n", addiction.StartedAt.Format("02.01.2006"), addiction.StartedAt.Format("15:04"))

		fmt.Printf("Streak:\t\t%d 🔥\n", streak)
		fmt.Printf("Goal:\t\t%d 🔥\n\n", addiction.StreakGoal)

		// if streak goal is reached, print message to ask the user if they want to increase the streak goal or not
		if streak >= addiction.StreakGoal {
			var daysStr = "Day"
			if streak != 1 {
				daysStr += "s"
			}

			fmt.Printf("Streak goal completed! %d/%d %s. Do you want to increase the streak goal? (yes/no)\n", streak, addiction.StreakGoal, daysStr)

			// get input from user
			var response string
			fmt.Print("Response (yes/no): ")
			fmt.Scanln(&response)
			if strings.ToLower(response) == "yes" {
				// ask user for new streak goal
				var newStreakGoalStr string

				fmt.Print("Enter new streak goal: ")
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

		fmt.Println()
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

	var daysStr = "Day"
	if streak != 1 {
		daysStr += "s"
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
