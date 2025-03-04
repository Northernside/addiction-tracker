package commands

import "fmt"

func Remove(args []string) error {
	if len(args) < 1 {
		fmt.Println("Usage: remove <addiction name>")
		return ErrNotEnoughArgs
	}

	name := args[0]
	err := RemoveAddiction(name)
	if err != nil {
		return err
	}

	fmt.Printf("Addiction '%s' removed\n", name)
	return nil
}
