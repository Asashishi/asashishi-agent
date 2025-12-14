package test

import (
	"asashishi-agent/global"
	"asashishi-agent/tools"
	"fmt"
)

func RunTest() {

	// TimeOps
	if tools.GetFormatedTime() == global.EmptyString {
		panic(ExceptionAtGetFormatedTime)
	}

	// FileOps
	if len(tools.GetFileList(global.EmptyString)) < 1 {
		panic(ExceptionAtGetFlieList)
	} else if !tools.CreateDir(CreateDirParam) {
		panic(ExceptionAtCreateDir)
	} else if !tools.CreateFile(CreateFileParam) {
		panic(WrontAtCreateFile)
	} else if !tools.MoveContent(CreateFileParam, MoveContentParam) {
		panic(ExceptionAtMoveContent)
	} else if !tools.AppendContentAtTail(MoveContentParam, AppendContentAtTailParam) {
		panic(ExceptionAtAppendContentAtTail)
	} else if tools.ReadFileContent(MoveContentParam) == global.EmptyString {
		panic(ExceptionAtReadFileContent)
	} else if !tools.RenewFileCache(MoveContentParam) {
		panic(ExceptionAtRenewFileCache)
	} else if !tools.FileContentRollBack(MoveContentParam) {
		panic(ExceptionAtFileContentRollBack)
	} else if position := tools.SearchFileContent(MoveContentParam, SearchFileContentParam); len(position) == 0 {
		panic(ExceptionAtSearchFileContent)
	} else if !tools.ReplaceFileContentByPosition(MoveContentParam, position[0], SearchFileContentParam) {
		panic(ExceptionAtReplaceFileContentByPosition)
	} else if !tools.DeleteFileContentByPosition(MoveContentParam, position[0]) {
		panic(ExceptionAtDeleteFileContentByPosition)
	} else if !tools.DeleteFileContent(MoveContentParam) {
		panic(ExceptionAtDeleteFileContent)
	} else if !tools.RemoveFile(MoveContentParam) {
		panic(ExceptionAtRemoveFile)
	} else if !tools.RemoveDir(CreateDirParam) {
		panic(ExceptionAtRemoveDir)
	}

	// ShellOps
	// 此处不单独测试跑命令的交互环境
	if !tools.AddCommands(AddCommandsParam) {
		panic(ExceptionAtAddCommands)
	} else if len(tools.GetCommands()) < 1 {
		panic(ExceptionAtGetCommands)
	} else if tools.AddCommands(AddCommandsParam) && !tools.PopCommands(1) && len(tools.GetCommands()) == 1 {
		panic(ExceptionAtPopCommands)
	} else if !tools.ClearCommands() {
		panic(ExceptionAtClearCommands)
	}

	// NetOps
	if len(tools.WebContentSearch(WebContentSearchParam)) < 1 {
		panic(ExceptionAtWebContentSearch)
	}

	fmt.Println(CompleteComment)
}
