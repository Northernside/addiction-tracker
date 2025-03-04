package commands

import "fmt"

func Reset(args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: reset <addiction name>")
		return ErrNotEnoughArgs
	}

	name := args[0]
	err := ResetAddiction(name)
	if err != nil {
		return err
	}

	fmt.Printf("Addiction '%s' reset\n", name)
	return nil
}
