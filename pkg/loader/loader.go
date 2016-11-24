package loader

import "io"

type Loader struct {
	source io.Reader // use bufio.Scanner to read by line
}
