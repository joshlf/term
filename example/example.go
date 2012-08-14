// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/joshlf13/term"
	"os"
)

func main() {
	// On windows, it is possible, though unlikely,
	// to fail to write to the console. An error
	// is returned.
	f := os.Stdout
	
	if err := term.Black(f, "Black"); err != nil {
		fmt.Printf("Stupid windows system calls: %v\n", err)
	}

	fmt.Println()
	term.White(f, "White")
	fmt.Println()
	term.Red(f, "Red")
	fmt.Println()
	term.LightRed(f, "Light Red")
	fmt.Println()
	term.Green(f, "Green")
	fmt.Println()
	term.LightGreen(f, "Light Green")
	fmt.Println()
	term.DarkYellow(f, "Dark Yellow")
	fmt.Println()
	term.LightYellow(f, "Light Yellow")
	fmt.Println()
	term.Blue(f, "Blue")
	fmt.Println()
	term.LightBlue(f, "Light Blue")
	fmt.Println()
	term.Magenta(f, "Magenta")
	fmt.Println()
	term.LightMagenta(f, "Light Magenta")
	fmt.Println()
	term.Cyan(f, "Cyan")
	fmt.Println()
	term.LightCyan(f, "Light Cyan")
	fmt.Println()
}
