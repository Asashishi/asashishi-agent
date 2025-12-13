package test

import (
	"asashishi-agent/global"
	"asashishi-agent/tools"
	"fmt"
)

func RunTest() {

	// TimeOps
	if tools.GetFormatedTime() == global.EmptyString {
		panic(WrongAtGetFormatedTime)
	}

	// FileOps
	if len(tools.GetFileList(global.EmptyString)) < 1 {
		panic(WrongAtGetFlieList)
	} else if !tools.CreateDir(CreateDirParam) {
		panic(WrongAtCreateDir)
	} else if !tools.CreateFile(CreateFileParam) {
		panic(WrontAtCreateFile)
	} else if !tools.MoveContent(CreateFileParam, MoveContentParam) {
		panic(WrongAtMoveContent)
	} else if !tools.AppendContentAtTail(MoveContentParam, AppendContentAtTailParam) {
		panic(WrongAtAppendContentAtTail)
	} else if tools.ReadFileContent(MoveContentParam) == global.EmptyString {
		panic(WrongAtReadFileContent)
	} else if !tools.RenewFileCache(MoveContentParam) {
		panic(WrongAtRenewFileCache)
	} else if !tools.FileContentRollBack(MoveContentParam) {
		panic(WrongAtFileContentRollBack)
	} else if position := tools.SearchFileContent(MoveContentParam, SearchFileContentParam); len(position) == 0 {
		panic(WrongAtSearchFileContent)
	} else if !tools.ReplaceFileContentByPosition(MoveContentParam, position[0], SearchFileContentParam) {
		panic(WrongAtReplaceFileContentByPosition)
	} else if !tools.DeleteFileContentByPosition(MoveContentParam, position[0]) {
		panic(WrongAtDeleteFileContentByPosition)
	} else if !tools.DeleteFileContent(MoveContentParam) {
		panic(WrongAtDeleteFileContent)
	} else if !tools.RemoveFile(MoveContentParam) {
		panic(WrongAtRemoveFile)
	} else if !tools.RemoveDir(CreateDirParam) {
		panic(WrongAtRemoveDir)
	}

	// ShellOps
	// 此处不单独测试跑命令的交互环境
	if !tools.AddCommands(AddCommandsParam) {
		panic(WrongAtAddCommands)
	} else if len(tools.GetCommands()) < 1 {
		panic(WrongAtGetCommands)
	} else if tools.AddCommands(AddCommandsParam) && !tools.PopCommands(1) && len(tools.GetCommands()) == 1 {
		panic(WrongAtPopCommands)
	} else if !tools.ClearCommands() {
		panic(WrongAtClearCommands)
	}

	// NetOps
	if len(tools.WebContentSearch(WebContentSearchParam)) < 1 {
		panic(WrongAtWebContentSearch)
	}

	fmt.Println(CompleteComment)
}
