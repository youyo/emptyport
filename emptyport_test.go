package emptyport_test

import (
	"github.com/youyo/emptyport"
	"testing"
)

func TestGet(t *testing.T) {
	port := emptyport.Get()
	if port > 1024 && port < 65535 {
		t.Log("ok")
	} else {
		t.Error("failed")
	}
}
