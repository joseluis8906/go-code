package logging

import (
	"log"
	"os"
)

func New() *log.Logger {
	l := log.New(os.Stderr, "[app:delivery] ", log.LstdFlags)
	l.SetFlags(log.LstdFlags | log.Llongfile)

	return l
}
