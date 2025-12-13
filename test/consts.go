package test

// TimeOps
const WrongAtGetFormatedTime string = "GetFormatedTime"

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

const WrongAtCreateDir string = "CreateDir"
const WrontAtCreateFile string = "CreateFile"
const WrongAtRemoveFile string = "RemoveFile"
const WrongAtRemoveDir string = "RemoveDir"
const WrongAtGetFlieList string = "GetFileList"
const WrongAtMoveContent string = "MoveContent"
const WrongAtRenewFileCache string = "RenewFileCache"
const WrongAtAppendContentAtTail string = "AppendContentAtTail"
const WrongAtSearchFileContent string = "SearchFileContent"
const WrongAtReadFileContent string = "ReadFileConten"
const WrongAtFileContentRollBack string = "FileContentRollBack"
const WrongAtDeleteFileContent string = "DeleteFileContent"
const WrongAtReplaceFileContentByPosition string = "ReplaceFileContentByPosition"
const WrongAtDeleteFileContentByPosition string = "DeleteFileContentByPosition"

// ShellOps
const AddCommandsParam string = "echo Asashishi"

const WrongAtAddCommands string = "AddCommands"
const WrongAtGetCommands string = "GetCommands"
const WrongAtPopCommands string = "PopCommands"
const WrongAtClearCommands string = "ClearCommands"

// NetOps
const WebContentSearchParam string = "https://github.com/Asashishi/asashishi-agent"

const WrongAtWebContentSearch string = "WebContentSearch"

// Complete
const CompleteComment string = "\nAll Test Past!"
