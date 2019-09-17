package main

import (
	"fmt"
	"os"

	"github.com/vysheradugi4/go-cli/actions"
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

	mainCommand := os.Args[1]

	err := actions.Execute(mainCommand)
	if err != nil {
		os.Exit(1)
	}
}
