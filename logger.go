package filelog

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/blbgo/general"
)

type fileLog struct {
	*os.File
}

// New provides a general.Logger interface that uses a file as storage
func New(config Config) (general.Logger, error) {
	return new(config, "log")
}

func new(config Config, name string) (general.Logger, error) {
	file, err := os.Create(
		filepath.Join(
			config.LogDir(),
			fmt.Sprintf("%v%v.txt", name, time.Now().Format("2006-01-02T15-04-05_000")),
		),
	)
	if err != nil {
		return nil, err
	}
	return &fileLog{File: file}, nil
}

func (r *fileLog) Log(v ...interface{}) error {
	r.WriteString(fmt.Sprint(v...))
	r.WriteString("\n")
	return nil
}

func (r *fileLog) Logf(format string, v ...interface{}) error {
	r.WriteString(fmt.Sprintf(format, v...))
	r.WriteString("\n")
	return nil
}

func (r *fileLog) Close() error {
	return r.File.Close()
}
