package bzip

import (
	"bytes"
	"io"
	"os/exec"
	"sync"
)

type writer struct {
	cmd     *exec.Cmd
	input   io.Writer
	writeMu sync.Mutex
}

func NewWriter(out io.Writer) io.WriteCloser {
	c := exec.Command("bzip2")
	stdin := &bytes.Buffer{}
	c.Stdin = stdin
	c.Stdout = out

	w := &writer{cmd: c, input: stdin}
	return w
}

func (w *writer) Write(data []byte) (int, error) {
	w.writeMu.Lock()
	defer w.writeMu.Unlock()
	return w.input.Write(data)
}

func (w *writer) Close() error {
	err := w.cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
