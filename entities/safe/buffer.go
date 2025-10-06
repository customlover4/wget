package safe

import (
	"bytes"
	"sync"
)

type Buffer struct {
	buf *bytes.Buffer
	mu  *sync.Mutex
}

func NewBuffer() *Buffer {
	return &Buffer{
		buf: new(bytes.Buffer),
		mu:  &sync.Mutex{},
	}
}

func (sb *Buffer) Len() int {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return sb.buf.Len()
}

func (sb *Buffer) Write(p []byte) (n int, err error) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return sb.buf.Write(p)
}

func (sb *Buffer) Buf() *bytes.Buffer {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return bytes.NewBuffer(sb.buf.Bytes())
}
