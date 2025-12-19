package global

import "testing"

func TestWaitNextFrame(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("WaitNextFrame() panic: %v", r)
		}
	}()
	WaitNextFrame(1000)
}
