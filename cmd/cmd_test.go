package cmd

import (
	"asashishi-agent/conf"
	"testing"
)

func TestBackupFiles(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("CmdTools() panic: %v", r)
		}
	}()
	CmdTools["-cls"]()
	if conf.EnvDetect() == conf.Windows {
		CmdTools["-rfile"]("cmd -rfile .\\init-cmd.go")
	} else {
		CmdTools["-rfile"]("cmd -rfile ./init-cmd.go")
	}
}
