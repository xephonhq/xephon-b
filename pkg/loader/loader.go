package loader

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/xephonhq/xephon-b/pkg/config"
)

type Loader struct {
	config config.LoaderConfig
	source io.Reader // use bufio.Scanner to read by line
}

func NewLoader(c config.LoaderConfig) *Loader {
	l := &Loader{}
	l.source = c.Source
	l.config = c
	return l
}

func (l *Loader) Start() {
	scanner := bufio.NewScanner(l.source)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
