package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("    ██████╗  ██████╗        ██████╗██╗     ██╗")
	fmt.Println("   ██╔════╝ ██╔═══██╗      ██╔════╝██║     ██║")
	fmt.Println("   ██║  ███╗██║   ██║█████╗██║     ██║     ██║")
	fmt.Println("   ██║   ██║██║   ██║╚════╝██║     ██║     ██║")
	fmt.Println("   ╚██████╔╝╚██████╔╝      ╚██████╗███████╗██║")
	fmt.Println("    ╚═════╝  ╚═════╝        ╚═════╝╚══════╝╚═╝")
	fmt.Println("")

	if len(os.Args) == 1 {
		fmt.Println("   ARGUMENTS REQUIRED!")
		fmt.Println("")
		fmt.Println("")
		os.Exit(0)
	}

	// mainCommand := os.Args[1]

	// err := actions.execute(mainCommand)
	// if err != nil {
	// 	os.Exit(1)
	// }
}
