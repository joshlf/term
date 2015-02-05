package term

import (
	"github.com/golang/crypto/terminal"
	"os"
)

// IsTerm returns true if os.Stdout is a terminal.
func IsTerm() bool {
	return FileIsTerm(os.Stdout)
}

// FileIsTerm returns true if f is a terminal.
func FileIsTerm(f *os.File) bool {
	return terminal.IsTerm(f.Fd())
}
