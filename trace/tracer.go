package trace

import "io"
import "fmt"

// Tracer is the inferface that describes an object capable of
// tracing events throughout code.
type Tracer interface {
	Trace(...interface{})
}

// New ...
func New(w io.Writer) Tracer {
	return &defaultTracer{out: w}
}

type defaultTracer struct {
	out io.Writer
}

func (t *defaultTracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
