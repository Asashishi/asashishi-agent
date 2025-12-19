package conf

import (
	"testing"
)

func TestEnvDetect(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("EnvDetect() panic: %v", r)
		}
	}()
	EnvDetect()
}
