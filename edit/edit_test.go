package edit

import (
	"os"
	"testing"
)

type editTest struct {
	command       string
	input, output string
}

var editTests = []editTest{
	{"touch", "foo", "foo"},
	{"bash", "#!/bin/bash\necho -n foo > $0", "foo"},
}

// TODO(synful): Make sure this path is actually illegal,
// not just unlikely to occur in practice.
var illegalPath = "/illegal/\n/path"

func TestEdit(t *testing.T) {
	originalVisual := os.Getenv("VISUAL")
	originalEditor := os.Getenv("EDITOR")
	defer func() {
		os.Setenv("VISUAL", originalVisual)
		os.Setenv("EDITOR", originalEditor)
	}()
	for _, test := range editTests {
		os.Setenv("VISUAL", test.command)
		out, err := Edit([]byte(test.input))
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		} else if string(out) != test.output {
			t.Errorf("unexpected output; expected %v; got %v", test.output, out)
		}
	}

	os.Setenv("VISUAL", illegalPath)
	out, err := Edit([]byte{})
	if err == nil {
		t.Errorf("unexpected nil error")
	}
	if out != nil {
		t.Errorf("unexpected output buffer; expected nil; got %v", out)
	}
}
