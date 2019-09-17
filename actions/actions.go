package actions

import (
	"fmt"
)

func execute(action string) error {
	switch action {
	case "new":
		fmt.Println("New selected")

	case "start":
		fmt.Println("Start selected")

	case "help":
		fmt.Println("Help selected")

	}

	return nil
}
