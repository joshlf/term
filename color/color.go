// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The term package allows easy handling of colors and terminal size cross-platform.
// It is compatible with unix variants and win32 terminals, but is incompatible with MS-DOS.
//
// Each function writes colored text directly to given file handle. If there was an error
// in printing, it returns that error, else nil. On unix systems, this will always be nil.
// On windows, however, it is possible for certain system calls to fail, and thus an
// error is possible.
package term

import "os"

func Black(f *os.File, s string) error {
	return black(f, s)
}

func White(f *os.File, s string) error {
	return white(f, s)
}

func Red(f *os.File, s string) error {
	return red(f, s)
}

// On non-windows systems,
// synonymous with Red
func LightRed(f *os.File, s string) error {
	return lightRed(f, s)
}

func Green(f *os.File, s string) error {
	return green(f, s)
}

// On non-windows systems,
// synonymous with Green
func LightGreen(f *os.File, s string) error {
	return lightGreen(f, s)
}

func DarkYellow(f *os.File, s string) error {
	return darkYellow(f, s)
}

// On non-windows systems,
// synonymous with DarkYellow
func LightYellow(f *os.File, s string) error {
	return lightYellow(f, s)
}

func Blue(f *os.File, s string) error {
	return blue(f, s)
}

// On non-windows systems,
// synonymous with Blue
func LightBlue(f *os.File, s string) error {
	return lightBlue(f, s)
}

func Magenta(f *os.File, s string) error {
	return magenta(f, s)
}

// On non-windows systems,
// synonymous with Magenta
func LightMagenta(f *os.File, s string) error {
	return lightMagenta(f, s)
}

func Cyan(f *os.File, s string) error {
	return cyan(f, s)
}

// On non-windows systems,
// synonymous with Cyan
func LightCyan(f *os.File, s string) error {
	return lightCyan(f, s)
}
