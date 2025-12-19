package tools

// test params mock

// TimeOps
const ExceptionAtGetFormatedTime string = "GetFormatedTime"

// FileOps
const CreateDirParamWindows string = `.\temp`
const CreateFileParamWindows string = `.\temp\test.txt`
const MoveContentParamWindows string = `.\temp\test.md`

const CreateDirParamLinux string = `./temp`
const CreateFileParamLinux string = `./temp/test.txt`
const MoveContentParamLinux string = `./temp/test.md`

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
const HttpSearchParam string = "https://github.com/Asashishi/asashishi-agent"

const ExceptionAtHttpSearch string = "HttpSearch"

const CompleteComment string = "✓ All Tool Test Past!"
