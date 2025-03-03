package commands

import (
	"fmt"
	"strings"
	"time"
)

func Live(args []string) error {
	// usage: live

	// load all addictions
	addictions, err := LoadAddictions()
	if err != nil {
		return err
	}

	if len(addictions) == 0 {
		fmt.Println("Keine Sucht vorhanden")
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
				streakDays := int(duration.Hours() / 24)
				streakHours := int(duration.Hours()) % 24
				streakMinutes := int(duration.Minutes()) % 60
				streakSeconds := int(duration.Seconds()) % 60

				fmt.Printf("Sucht:\t\t%s\n", addiction.Name)
				fmt.Printf("Gestartet am:\t%s um %s\n\n", addiction.StartedAt.Format("02.01.2006"), addiction.StartedAt.Format("15:04"))

				var daysStr string = "Tag"
				if streakDays != 1 {
					daysStr += "e"
				}

				fmt.Printf("Streak:\t\t%d %s 🔥\n", streakDays, daysStr)
				fmt.Printf("Ziel:\t\t%d %s 🔥\n\n", addiction.StreakGoal, daysStr)

				// print progress bars for days, hours, minutes, and seconds
				fmt.Printf("Tage:\t\t%s", liveProgress(streakDays, addiction.StreakGoal))
				fmt.Printf("Stunden:\t%s", liveProgress(streakHours, 24))
				fmt.Printf("Minuten:\t%s", liveProgress(streakMinutes, 60))
				fmt.Printf("Sekunden:\t%s", liveProgress(streakSeconds, 60))

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
