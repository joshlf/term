package term

import (
	"fmt"
	t "github.com/mewkiz/pkg/term"
)

func black(s string) error {
	fmt.Print(t.Color(s, t.FgBlack))
	return nil
}

func white(s string) error {
	fmt.Print(t.Color(s, t.FgWhite))
	return nil
}

func red(s string) error {
	fmt.Print(t.Color(s, t.FgRed))
	return nil
}

func lightRed(s string) error {
	fmt.Print(t.Color(s, t.FgRed))
	return nil
}

func green(s string) error {
	fmt.Print(t.Color(s, t.FgGreen))
	return nil
}

func lightGreen(s string) error {
	fmt.Print(t.Color(s, t.FgGreen))
	return nil
}

func darkYellow(s string) error {
	fmt.Print(t.Color(s, t.FgYellow))
	return nil
}

func lightYellow(s string) error {
	fmt.Print(t.Color(s, t.FgYellow))
	return nil
}

func blue(s string) error {
	fmt.Print(t.Color(s, t.FgBlue))
	return nil
}

func lightBlue(s string) error {
	fmt.Print(t.Color(s, t.FgBlue))
	return nil
}

func magenta(s string) error {
	fmt.Print(t.Color(s, t.FgMagenta))
	return nil
}

func lightMagenta(s string) error {
	fmt.Print(t.Color(s, t.FgMagenta))
	return nil
}

func cyan(s string) error {
	fmt.Print(t.Color(s, t.FgCyan))
	return nil
}

func lightCyan(s string) error {
	fmt.Print(t.Color(s, t.FgCyan))
	return nil
}
