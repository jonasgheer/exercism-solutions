package paasio

import (
	"io"
	"sync"
)

type writeCounter struct {
	sync.RWMutex
	writer       io.Writer
	writtenBytes int64
	callCount    int
}

type readCounter struct {
	sync.RWMutex
	reader    io.Reader
	readBytes int64
	callCount int
}

type readWriteCounter struct {
	ReadCounter
	WriteCounter
}

// Define readCounter and writeCounter types here.

// For the return of the function NewReadWriteCounter, you must also define a type that satisfies the ReadWriteCounter interface.

func NewWriteCounter(writer io.Writer) WriteCounter {
	return &writeCounter{writer: writer}
}

func NewReadCounter(reader io.Reader) ReadCounter {
	return &readCounter{reader: reader}
}

func NewReadWriteCounter(readwriter io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{NewReadCounter(readwriter), NewWriteCounter(readwriter)}
}

func (rc *readCounter) Read(p []byte) (int, error) {
	rc.Lock()
	defer rc.Unlock()
	count, err := rc.reader.Read(p)
	rc.readBytes += int64(count)
	rc.callCount++
	return count, err
}

func (rc *readCounter) ReadCount() (int64, int) {
	rc.Lock()
	defer rc.Unlock()
	return rc.readBytes, rc.callCount
}

func (wc *writeCounter) Write(p []byte) (int, error) {
	wc.Lock()
	defer wc.Unlock()
	count, err := wc.writer.Write(p)
	wc.writtenBytes += int64(count)
	wc.callCount++
	return count, err
}

func (wc *writeCounter) WriteCount() (int64, int) {
	wc.Lock()
	defer wc.Unlock()
	return wc.writtenBytes, wc.callCount
}
