package actions

import (
	"fmt"
)

// Execute ...
func Execute(action string) error {
	switch action {
	case "new":
		fmt.Println("  >   Web service (RESTful API).")
		fmt.Println("      Web service (RESTful API + WebSocket).")
		fmt.Println("      Web service (GRPC).")
		fmt.Println("      Web service (GraphQL).")
		fmt.Println("      MicroService (HTTP).")
		fmt.Scanln()

	case "start":
		fmt.Println("Start selected")

	case "help":
		fmt.Println("Help selected")

	}

	return nil
}
