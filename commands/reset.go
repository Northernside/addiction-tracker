package commands

import "fmt"

func Reset(args []string) error {
	// usage: reset <addiction name>

	if len(args) < 1 {
		fmt.Println("Benutzung: reset <addiction name>")
		return ErrNotEnoughArgs
	}

	name := args[0]
	err := ResetAddiction(name)
	if err != nil {
		return err
	}

	fmt.Printf("Sucht %s wurde zurückgesetzt!\n", name)
	return nil
}
