package test

// TimeOps
const ExceptionAtGetFormatedTime string = "GetFormatedTime"

// FileOps
const CreateDirParam string = `.\temp`
const CreateFileParam string = `.\temp\test.txt`
const MoveContentParam string = `.\temp\test.md`
const AppendContentAtTailParam string = `
1: 测试文本
2: 测试文本
3: 测试文本
`
const SearchFileContentParam string = "1: 测试文本"

const ExceptionAtCreateDir string = "CreateDir"
const WrontAtCreateFile string = "CreateFile"
const ExceptionAtRemoveFile string = "RemoveFile"
const ExceptionAtRemoveDir string = "RemoveDir"
const ExceptionAtGetFlieList string = "GetFileList"
const ExceptionAtMoveContent string = "MoveContent"
const ExceptionAtRenewFileCache string = "RenewFileCache"
const ExceptionAtAppendContentAtTail string = "AppendContentAtTail"
const ExceptionAtSearchFileContent string = "SearchFileContent"
const ExceptionAtReadFileContent string = "ReadFileConten"
const ExceptionAtFileContentRollBack string = "FileContentRollBack"
const ExceptionAtDeleteFileContent string = "DeleteFileContent"
const ExceptionAtReplaceFileContentByPosition string = "ReplaceFileContentByPosition"
const ExceptionAtDeleteFileContentByPosition string = "DeleteFileContentByPosition"

// ShellOps
const AddCommandsParam string = "echo Asashishi"

const ExceptionAtAddCommands string = "AddCommands"
const ExceptionAtGetCommands string = "GetCommands"
const ExceptionAtPopCommands string = "PopCommands"
const ExceptionAtClearCommands string = "ClearCommands"

// NetOps
const WebContentSearchParam string = "https://github.com/Asashishi/asashishi-agent"

const ExceptionAtWebContentSearch string = "WebContentSearch"

const CompleteComment string = "✓  All Test Past!"
