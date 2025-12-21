package backup

import "asashishi-agent/global"

var StartBackupComment string = global.GetStyledSystemComent("-- Start Backup")
var BackupCompletedComment string = global.GetStyledSystemComent("All Completed! --")
var NoFileToBackupComment string = global.GetStyledSystemComent("-- No files to backup --")
