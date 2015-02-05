// Package edit allows programs to ask for user input through a text editor
// in the manner of git commit or contrab -e.
package edit

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var defaultEditor = "vi"

// Edit is shorthand for EditWith(Editor(), contents).
func Edit(contents []byte) ([]byte, error) {
	return EditWith(Editor(), contents)
}

// EditWith opens the specified editor populated with
// contents for the uesr to edit, and returns the resulting
// contents, including any changes made by the user.
// The returned byte slice will be non-nil if and
// only if the returned error is nil.
func EditWith(editor string, contents []byte) ([]byte, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, err
	}
	defer os.Remove(f.Name())
	_, err = f.Write(contents)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(editor, f.Name())
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%v %v: %v", editor, f.Name(), err)
	}
	_, err = f.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

// Editor returns the default editor. It first checks the
// VISUAL environment variable, then the EDITOR environment
// variable, and then falls back to vi on Unix or notepad.exe
// on Windows.
func Editor() string {
	env := os.Environ()
	editor, ok := findVar(env, "VISUAL")
	if ok {
		return editor
	}
	editor, ok = findVar(env, "EDITOR")
	if ok {
		return editor
	}
	return defaultEditor
}

func findVar(env []string, v string) (string, bool) {
	str := v + "="
	l := len(str)
	for _, e := range env {
		if len(e) >= l && e[:l] == str {
			return e[l:], true
		}
	}
	return "", false
}
