package poker

import (
	"io"
	"os"
)

type Tape struct {
	file *os.File
}

func (t *Tape) Write(data []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, io.SeekStart)
	return t.file.Write(data)
}
