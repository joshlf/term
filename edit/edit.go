// Package edit allows programs to ask for user input through a text editor
// in the manner of git commit or contrab -e.
package edit

import (
	"os"
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
	return editWith(editor, contents)
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

// EditorFallback behaves like Editor, except that it
// falls back to fallback instead of vi on Unix or
// notepad.exe on Windows.
func EditorFallback(fallback string) string {
	env := os.Environ()
	editor, ok := findVar(env, "VISUAL")
	if ok {
		return editor
	}
	editor, ok = findVar(env, "EDITOR")
	if ok {
		return editor
	}
	return fallback
}

// Use findVar because os.Getenv is not powerful enough:
// it cannot distinguish between an unset variable and
// one whose value is the empty string.
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
