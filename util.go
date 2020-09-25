package gomsf

import (
	"bytes"
	"sync"
)

var (
	bufpool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
)
