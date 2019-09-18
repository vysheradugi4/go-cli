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
	fmt.Println("")

	if len(os.Args) == 1 {
		fmt.Println(" Available Commands:")
		fmt.Println("")
		fmt.Println(" new    (n)  Creates a new workspace and an initial Golang app.")
		fmt.Println(" add         Adds support for an external library to your project.")
		fmt.Println(" serve  (s)  Builds and serves your app, rebuilding on file changes.")
		fmt.Println(" test   (t)  Runs unit tests in a project.")
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
