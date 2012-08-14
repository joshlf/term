// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package term

import (
	"github.com/anschelsc/doscolor"
	"fmt"
	"os"
)

func black(s string) error {
	return text(s, []doscolor.Color{doscolor.Black})
}

func white(s string) error {
	return text(s, []doscolor.Color{doscolor.White, doscolor.Bright})
}

func red(s string) error {
	return text(s, []doscolor.Color{doscolor.Red})
}

func lightRed(s string) error {
	return text(s, []doscolor.Color{doscolor.Red, doscolor.Bright})
}

func green(s string) error {
	return text(s, []doscolor.Color{doscolor.Green})
}

func lightGreen(s string) error {
	return text(s, []doscolor.Color{doscolor.Green, doscolor.Bright})
}

func darkYellow(s string) error {
	return text(s, []doscolor.Color{doscolor.Yellow})
}

func lightYellow(s string) error {
	return text(s, []doscolor.Color{doscolor.Yellow, doscolor.Bright})
}

func blue(s string) error {
	return text(s, []doscolor.Color{doscolor.Blue})
}

func lightBlue(s string) error {
	return text(s, []doscolor.Color{doscolor.Blue, doscolor.Bright})
}

func magenta(s string) error {
	return text(s, []doscolor.Color{doscolor.Magenta})
}

func lightMagenta(s string) error {
	return text(s, []doscolor.Color{doscolor.Magenta, doscolor.Bright})
}

func cyan(s string) error {
	return text(s, []doscolor.Color{doscolor.Cyan})
}

func lightCyan(s string) error {
	return text(s, []doscolor.Color{doscolor.Cyan, doscolor.Bright})
}

// Return an error if the text wasn't printed;
// otherwise return nil
func text(s string, c []doscolor.Color) error {
	output := doscolor.NewWrapper(os.Stdout)
	if err := output.Save(); err != nil {
		return err
	}
	if len(c) < 1 {
		fmt.Fprintln(output, s)
	}
	
	var cbits doscolor.Color = 0
	for _, color := range c {
		cbits |= color
	}
	
	if err := output.Set(cbits); err != nil {
		return err
	}
	fmt.Fprint(output, s)
	
	// Don't check for an error here, 
	// because text has been printed
	output.Restore()
	return nil
}