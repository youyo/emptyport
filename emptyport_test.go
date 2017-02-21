package emptyport_test

import (
	"testing"

	"github.com/youyo/emptyport"
)

func TestGet(t *testing.T) {
	port, err := emptyport.Get()
	if port > 1024 && port < 65535 && err == nil {
		t.Log("ok")
	} else {
		t.Error("failed")
	}
}
