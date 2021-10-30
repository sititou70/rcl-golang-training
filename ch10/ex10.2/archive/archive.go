package archive

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
)

type Format struct {
	Signature []byte
	Offset    int
	FileNames func(f *os.File) ([]string, error)
}

var (
	formatsMu     sync.Mutex
	atomicFormats atomic.Value
)

func RegisterFormat(format Format) {
	formatsMu.Lock()
	formats, _ := atomicFormats.Load().([]Format)
	atomicFormats.Store(append(formats, format))
	formatsMu.Unlock()
}

func FileNames(f *os.File) ([]string, error) {
	formats, _ := atomicFormats.Load().([]Format)
	for _, format := range formats {
		sig := make([]byte, len(format.Signature))
		f.ReadAt(sig, int64(format.Offset))
		if bytes.Equal(sig, format.Signature) {
			return format.FileNames(f)
		}
	}

	return nil, fmt.Errorf("unknown format")
}
