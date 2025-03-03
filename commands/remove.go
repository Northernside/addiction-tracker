package commands

import "fmt"

func Remove(args []string) error {
	// usage: remove <addiction name>

	if len(args) < 1 {
		fmt.Println("Benutzung: remove <addiction name>")
		return ErrNotEnoughArgs
	}

	name := args[0]
	err := RemoveAddiction(name)
	if err != nil {
		return err
	}

	fmt.Printf("Sucht %s wurde entfernt!\n", name)
	return nil
}
