// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The term package allows easy handling of colors and terminal size cross-platform.
// It is compatible with unix variants and win32 terminals, but is incompatible with MS-DOS.
//
// Each function prints colored text directly to the console. If there was an error
// in printing, it returns that error, else nil. On unix systems, this will always be nil.
// On windows, however, it is possible for certain system calls to fail, and thus an
// error is possible.
package term

func Black(s string) error {
	return black(s)
}

func White(s string) error {
	return white(s)
}

func Red(s string) error {
	return red(s)
}

// On non-windows systems,
// synonymous with Red
func LightRed(s string) error {
	return lightRed(s)
}

func Green(s string) error {
	return green(s)
}

// On non-windows systems,
// synonymous with Green
func LightGreen(s string) error {
	return lightGreen(s)
}

func DarkYellow(s string) error {
	return darkYellow(s)
}

// On non-windows systems,
// synonymous with DarkYellow
func LightYellow(s string) error {
	return lightYellow(s)
}

func Blue(s string) error {
	return blue(s)
}

// On non-windows systems,
// synonymous with Blue
func LightBlue(s string) error {
	return lightBlue(s)
}

func Magenta(s string) error {
	return magenta(s)
}

// On non-windows systems,
// synonymous with Magenta
func LightMagenta(s string) error {
	return lightMagenta(s)
}

func Cyan(s string) error {
	return cyan(s)
}

// On non-windows systems,
// synonymous with Cyan
func LightCyan(s string) error {
	return lightCyan(s)
}
