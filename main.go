package main

import (
	"fmt"
	"os"

	"github.com/vysheradugi4/go-cli/actions"
	"golang.org/x/sys/windows"
)

const (
	// GoColor Cian color for keys interface elements
	GoColor = "\033[1;36m%s\033[0m"

	// NoticeColor Deep marine color
	NoticeColor = "\033[1;36m%s\033[0m"

	// ErrorColor Error messages color, brightly red.
	ErrorColor = "\033[1;31m%s\033[0m"
)

func init() {
	stdout := windows.Handle(os.Stdout.Fd())
	var originalMode uint32

	windows.GetConsoleMode(stdout, &originalMode)
	windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}

func main() {
	fmt.Println()
	fmt.Println()
	fmt.Printf(GoColor, "    ██████╗  ██████╗        ██████╗██╗     ██╗")
	fmt.Println()
	fmt.Printf(GoColor, "   ██╔════╝ ██╔═══██╗      ██╔════╝██║     ██║")
	fmt.Println()
	fmt.Printf(GoColor, "   ██║  ███╗██║   ██║█████╗██║     ██║     ██║")
	fmt.Println()
	fmt.Printf(GoColor, "   ██║   ██║██║   ██║╚════╝██║     ██║     ██║")
	fmt.Println()
	fmt.Printf(GoColor, "   ╚██████╔╝╚██████╔╝      ╚██████╗███████╗██║")
	fmt.Println()
	fmt.Printf(GoColor, "    ╚═════╝  ╚═════╝        ╚═════╝╚══════╝╚═╝")
	fmt.Println()
	fmt.Println()

	if len(os.Args) == 1 {
		fmt.Println(" Available Commands:")
		fmt.Println()

		fmt.Printf(NoticeColor, " new    (n)  ")
		fmt.Println("Creates a new workspace and an initial Golang app.")

		fmt.Printf(NoticeColor, " add         ")
		fmt.Println("Adds support for an external library to your project.")

		fmt.Printf(NoticeColor, " serve  (s)  ")
		fmt.Println("Builds and serves your app, rebuilding on file changes.")

		fmt.Printf(NoticeColor, " test   (t)  ")
		fmt.Println("Runs unit tests in a project.")

		fmt.Println()
		fmt.Println()
		os.Exit(0)
	}

	mainCommand := os.Args[1]

	err := actions.Execute(mainCommand)
	if err != nil {
		os.Exit(1)
	}
}
