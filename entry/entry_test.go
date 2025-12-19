package entry

import (
	"testing"
)

func TestInitCli(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("InitCli() panic: %v", r)
		}
	}()
	InitCli()
}

func TestInitWeb(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("InitWeb() panic: %v", r)
		}
	}()
	InitWeb()
}
