// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package term

import (
	"fmt"
	t "github.com/mewkiz/pkg/term"
	"os"
)

func black(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgBlack))
	return nil
}

func white(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgWhite))
	return nil
}

func red(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgRed))
	return nil
}

func lightRed(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgRed))
	return nil
}

func green(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgGreen))
	return nil
}

func lightGreen(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgGreen))
	return nil
}

func darkYellow(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgYellow))
	return nil
}

func lightYellow(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgYellow))
	return nil
}

func blue(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgBlue))
	return nil
}

func lightBlue(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgBlue))
	return nil
}

func magenta(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgMagenta))
	return nil
}

func lightMagenta(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgMagenta))
	return nil
}

func cyan(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgCyan))
	return nil
}

func lightCyan(s string, f *os.File) error {
	fmt.Fprint(f, t.Color(s, t.FgCyan))
	return nil
}
