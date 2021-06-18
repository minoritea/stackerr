package stackerr_test

import (
	"github.com/minoritea/stackerr"
	"testing"
)

func TestNew(t *testing.T) {
	err := func() error {
		return stackerr.New("foo")
	}()
	stack := stackerr.RootStackTrace(err)
	if stack == nil {
		t.Fatal("stack trace is nil")
	}
	t.Logf("%+v", err)
}
