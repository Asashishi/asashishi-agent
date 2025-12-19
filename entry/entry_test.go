package entry

import (
	"testing"
	"time"
)

func TestInitCli(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("InitCli() panic: %v", r)
		}
	}()
	InitCli()
	time.Sleep(time.Second * 5)
}

func TestInitWeb(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("InitWeb() panic: %v", r)
		}
	}()
	InitWeb()
	time.Sleep(time.Second * 5)
}
