package backup

import (
	"asashishi-agent/conf"
	"testing"
)

func TestBackupFiles(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("BackupFiles() panic: %v", r)
		}
	}()
	BackupFiles(conf.EnvDetect())
}
