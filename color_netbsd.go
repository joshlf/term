// Copyright 2012 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package term

import (
	"fmt"
	t "github.com/mewkiz/pkg/term"
	"os"
)

func black(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgBlack))
	return nil
}

func white(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgWhite))
	return nil
}

func red(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgRed))
	return nil
}

func lightRed(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgRed))
	return nil
}

func green(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgGreen))
	return nil
}

func lightGreen(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgGreen))
	return nil
}

func darkYellow(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgYellow))
	return nil
}

func lightYellow(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgYellow))
	return nil
}

func blue(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgBlue))
	return nil
}

func lightBlue(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgBlue))
	return nil
}

func magenta(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgMagenta))
	return nil
}

func lightMagenta(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgMagenta))
	return nil
}

func cyan(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgCyan))
	return nil
}

func lightCyan(f *os.File, s string) error {
	fmt.Fprint(f, t.Color(s, t.FgCyan))
	return nil
}
