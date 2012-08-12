package main

import (
	"fmt"
	"github.com/joshlf13/term"
)

func main() {
	// On windows, it is possible, though unlikely,
	// to fail to write to the console. An error
	// is returned.
	if err := term.Black("Black"); err != nil {
		fmt.Printf("Stupid windows system calls: %v\n", err)
	}

	fmt.Println()
	term.White("White")
	fmt.Println()
	term.Red("Red")
	fmt.Println()
	term.LightRed("Light Red")
	fmt.Println()
	term.Green("Green")
	fmt.Println()
	term.LightGreen("Light Green")
	fmt.Println()
	term.DarkYellow("Dark Yellow")
	fmt.Println()
	term.LightYellow("Light Yellow")
	fmt.Println()
	term.Blue("Blue")
	fmt.Println()
	term.LightBlue("Light Blue")
	fmt.Println()
	term.Magenta("Magenta")
	fmt.Println()
	term.LightMagenta("Light Magenta")
	fmt.Println()
	term.Cyan("Cyan")
	fmt.Println()
	term.LightCyan("Light Cyan")
	fmt.Println()
}
