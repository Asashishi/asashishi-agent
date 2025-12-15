package backup

import (
	"asashishi-agent/global"
	"asashishi-agent/ui"
)

var NoFileToBackupComment string = global.LineBreakString + ui.WidthStyle(
	"-- No files to backup --",
	global.SystemCommentStyle,
)

var StartBackupComment string = global.LineBreakString + ui.WidthStyle(
	"-- Start Backup",
	global.SystemCommentStyle,
)

var BackupCompletedComment string = ui.WidthStyle(
	"All Completed! --",
	global.SystemCommentStyle,
)
