package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return for New should not be nil")
	} else {
		testStr := "Hello trace package."
		tracer.Trace(testStr)
		if buf.String() != (testStr + "\n") {
			t.Errorf("Trace should not write '%s'", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	silentTracer := Off()
	silentTracer.Trace("something")
}
