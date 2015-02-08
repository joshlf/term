// +build !windows

package edit

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func editWith(editor string, contents []byte) ([]byte, error) {
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
