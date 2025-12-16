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
		panic(global.GetStyledError(ExceptionAtGetFlieList))
	} else if !tools.CreateDir(CreateDirParam) {
		panic(global.GetStyledError(ExceptionAtCreateDir))
	} else if !tools.CreateFile(CreateFileParam) {
		panic(global.GetStyledError(WrontAtCreateFile))
	} else if !tools.MoveContent(CreateFileParam, MoveContentParam) {
		panic(global.GetStyledError(ExceptionAtMoveContent))
	} else if !tools.AppendContentAtTail(MoveContentParam, AppendContentAtTailParam) {
		panic(global.GetStyledError(ExceptionAtAppendContentAtTail))
	} else if tools.ReadFileContent(MoveContentParam) == global.EmptyString {
		panic(global.GetStyledError(ExceptionAtReadFileContent))
	} else if !tools.RenewFileCache(MoveContentParam) {
		panic(global.GetStyledError(ExceptionAtRenewFileCache))
	} else if !tools.FileContentRollBack(MoveContentParam) {
		panic(global.GetStyledError(ExceptionAtFileContentRollBack))
	} else if position := tools.SearchFileContent(MoveContentParam, SearchFileContentParam); len(position) == 0 {
		panic(global.GetStyledError(ExceptionAtSearchFileContent))
	} else if !tools.ReplaceFileContentByPosition(MoveContentParam, position[0], SearchFileContentParam) {
		panic(global.GetStyledError(ExceptionAtReplaceFileContentByPosition))
	} else if !tools.DeleteFileContentByPosition(MoveContentParam, position[0]) {
		panic(global.GetStyledError(ExceptionAtDeleteFileContentByPosition))
	} else if !tools.DeleteFileContent(MoveContentParam) {
		panic(global.GetStyledError(ExceptionAtDeleteFileContent))
	} else if !tools.RemoveFile(MoveContentParam) {
		panic(global.GetStyledError(ExceptionAtRemoveFile))
	} else if !tools.RemoveDir(CreateDirParam) {
		panic(global.GetStyledError(ExceptionAtRemoveDir))
	}

	// ShellOps
	// 此处不单独测试跑命令的交互环境
	if !tools.AddCommands(AddCommandsParam) {
		panic(global.GetStyledError(ExceptionAtAddCommands))
	} else if len(tools.GetCommands()) < 1 {
		panic(global.GetStyledError(ExceptionAtGetCommands))
	} else if tools.AddCommands(AddCommandsParam) && !tools.PopCommands(1) && len(tools.GetCommands()) == 1 {
		panic(global.GetStyledError(ExceptionAtPopCommands))
	} else if !tools.ClearCommands() {
		panic(global.GetStyledError(ExceptionAtClearCommands))
	}

	// NetOps
	if len(tools.WebContentSearch(WebContentSearchParam)) < 1 {
		panic(global.GetStyledError(ExceptionAtWebContentSearch))
	}

	fmt.Println(global.GetStyledSuccess(CompleteComment))
}
