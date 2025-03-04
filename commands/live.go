package commands

import (
	"fmt"
	"strings"
	"time"
)

func Live(args []string) error {
	addictions, err := LoadAddictions()
	if err != nil {
		return err
	}

	if len(addictions) == 0 {
		fmt.Println("No addictions found")
		return nil
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("\033[H\033[2J")
			for _, addiction := range addictions {
				// calculate the streak correctly
				duration := time.Since(addiction.StartedAt)
				streakYears := int(duration.Hours() / (24 * 365))
				streakDays := int(duration.Hours() / 24)
				streakHours := int(duration.Hours()) % 24
				streakMinutes := int(duration.Minutes()) % 60
				streakSeconds := int(duration.Seconds()) % 60

				fmt.Printf("Addiction:\t%s\n", addiction.Name)
				fmt.Printf("Started on:\t%s at %s\n\n", addiction.StartedAt.Format("02.01.2006"), addiction.StartedAt.Format("15:04"))

				if addiction.StreakGoal != -1 {
					var daysStr string = "Day"
					if streakDays != 1 {
						daysStr += "s"
					}

					fmt.Printf("Streak:\t\t%d %s 🔥\n", streakDays, daysStr)
					fmt.Printf("Goal:\t\t%d %s 🔥\n\n", addiction.StreakGoal, daysStr)
				}

				// print progress bars for days, hours, minutes, and seconds
				fmt.Printf("Years:\t\t%s", liveProgress(streakYears, 2))
				fmt.Printf("Days:\t\t%s", liveProgress(streakDays, 365))
				fmt.Printf("Hours:\t\t%s", liveProgress(streakHours, 24))
				fmt.Printf("Minutes:\t%s", liveProgress(streakMinutes, 60))
				fmt.Printf("Seconds:\t%s", liveProgress(streakSeconds, 60))

				fmt.Printf("\n\n")
			}
		}
	}
}

func liveProgress(current int, total int) string {
	const barWidth = 60

	blocks := int(float64(current) / float64(total) * barWidth)
	fullBlocks := blocks
	emptyBlocks := barWidth - blocks

	return fmt.Sprintf("[%s%s] %d/%d\n",
		strings.Repeat("█", fullBlocks),
		strings.Repeat("░", emptyBlocks),
		current, total)
}
