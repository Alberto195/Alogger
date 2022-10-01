package alog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

type AsyncLogger struct {
	mu      sync.Mutex
	out     io.Writer
	buf     []byte
	builder *strings.Builder
	prefix  string
	logCh   chan string
}

func NewAsyncLogger(prefix string, out io.Writer) *AsyncLogger {
	aLog := &AsyncLogger{
		out:     out,
		logCh:   make(chan string),
		builder: &strings.Builder{},
		prefix:  prefix,
	}
	aLog.printLog()
	return aLog
}

func (a *AsyncLogger) Info(s string) {
	a.logCh <- s
}

func (a *AsyncLogger) printLog() {
	go func() {
		defer close(a.logCh)
		for s := range a.logCh {
			a.mu.Lock()
			a.builder.Reset()
			a.builder.WriteString(a.prefix)
			a.builder.WriteString(time.Now().String())
			a.builder.WriteString(": ")
			a.buf = a.buf[:0]
			a.buf = append(a.buf, a.builder.String()...)
			a.buf = append(a.buf, s...)
			if len(s) == 0 || s[len(s)-1] != '\n' {
				a.buf = append(a.buf, '\n')
			}
			_, err := a.out.Write(a.buf)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v can't write a log: %v", time.Now().String(), err)
			}
			a.mu.Unlock()
		}
	}()
}
