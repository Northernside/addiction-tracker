package main

import (
	"fmt"
	"iamsober-tui/commands"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		err := commands.Help(args)
		if err != nil {
			panic(err)
		}

		return
	}

	for _, cmd := range commands.Commands {
		for _, key := range cmd.Keys {
			if key != args[0] {
				continue
			}

			var err error
			var args []string = []string{}

			if len(os.Args) > 2 {
				args = os.Args[2:]
			}

			err = cmd.Fn(args)
			if err != nil {
				fmt.Println(err)
				return
			}

			return
		}
	}
}
