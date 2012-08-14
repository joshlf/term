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

func Black(s string, f *os.File) error {
	return black(s, f)
}

func White(s string, f *os.File) error {
	return white(s, f)
}

func Red(s string, f *os.File) error {
	return red(s, f)
}

// On non-windows systems,
// synonymous with Red
func LightRed(s string, f *os.File) error {
	return lightRed(s, f)
}

func Green(s string, f *os.File) error {
	return green(s, f)
}

// On non-windows systems,
// synonymous with Green
func LightGreen(s string, f *os.File) error {
	return lightGreen(s, f)
}

func DarkYellow(s string, f *os.File) error {
	return darkYellow(s, f)
}

// On non-windows systems,
// synonymous with DarkYellow
func LightYellow(s string, f *os.File) error {
	return lightYellow(s, f)
}

func Blue(s string, f *os.File) error {
	return blue(s, f)
}

// On non-windows systems,
// synonymous with Blue
func LightBlue(s string, f *os.File) error {
	return lightBlue(s, f)
}

func Magenta(s string, f *os.File) error {
	return magenta(s, f)
}

// On non-windows systems,
// synonymous with Magenta
func LightMagenta(s string, f *os.File) error {
	return lightMagenta(s, f)
}

func Cyan(s string, f *os.File) error {
	return cyan(s, f)
}

// On non-windows systems,
// synonymous with Cyan
func LightCyan(s string, f *os.File) error {
	return lightCyan(s, f)
}
