package utils

import (
	"bytes"
	"log"
	"testing"
)

func TestHandleError_WithNilError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	HandleError(nil)

	if buf.String() != "" {
		t.Errorf("HandleError() = %q; want empty output", buf.String())
	}
}
