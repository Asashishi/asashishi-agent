package tools

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"fmt"
	"testing"
)

func ToolTest() {

	var env = conf.EnvDetect()

	// TimeOps
	if GetFormatedTime() == global.EmptyString {
		panic(ExceptionAtGetFormatedTime)
	}

	// FileOps
	if len(GetFileList(global.EmptyString)) < 1 {
		panic(global.GetStyledError(ExceptionAtGetFlieList))
	}
	if env == conf.Windows {
		if !CreateDir(CreateDirParamWindows) {
			panic(global.GetStyledError(ExceptionAtCreateDir))
		} else if !CreateFile(CreateFileParamWindows) {
			panic(global.GetStyledError(WrontAtCreateFile))
		} else if !MoveContent(CreateFileParamWindows, MoveContentParamWindows) {
			panic(global.GetStyledError(ExceptionAtMoveContent))
		} else if !AppendContentAtTail(MoveContentParamWindows, AppendContentAtTailParam) {
			panic(global.GetStyledError(ExceptionAtAppendContentAtTail))
		} else if ReadFileContent(MoveContentParamWindows) == global.EmptyString {
			panic(global.GetStyledError(ExceptionAtReadFileContent))
		} else if !RenewFileCache(MoveContentParamWindows) {
			panic(global.GetStyledError(ExceptionAtRenewFileCache))
		} else if !FileContentRollBack(MoveContentParamWindows) {
			panic(global.GetStyledError(ExceptionAtFileContentRollBack))
		} else if position := SearchFileContent(MoveContentParamWindows, SearchFileContentParam); len(position) == 0 {
			panic(global.GetStyledError(ExceptionAtSearchFileContent))
		} else if !ReplaceFileContentByPosition(MoveContentParamWindows, position[0], SearchFileContentParam) {
			panic(global.GetStyledError(ExceptionAtReplaceFileContentByPosition))
		} else if !DeleteFileContentByPosition(MoveContentParamWindows, position[0]) {
			panic(global.GetStyledError(ExceptionAtDeleteFileContentByPosition))
		} else if !DeleteFileContent(MoveContentParamWindows) {
			panic(global.GetStyledError(ExceptionAtDeleteFileContent))
		} else if !RemoveFile(MoveContentParamWindows) {
			panic(global.GetStyledError(ExceptionAtRemoveFile))
		} else if !RemoveDir(CreateDirParamWindows) {
			panic(global.GetStyledError(ExceptionAtRemoveDir))
		}
	} else {
		if !CreateDir(CreateDirParamLinux) {
			panic(global.GetStyledError(ExceptionAtCreateDir))
		} else if !CreateFile(CreateFileParamLinux) {
			panic(global.GetStyledError(WrontAtCreateFile))
		} else if !MoveContent(CreateFileParamLinux, MoveContentParamLinux) {
			panic(global.GetStyledError(ExceptionAtMoveContent))
		} else if !AppendContentAtTail(MoveContentParamLinux, AppendContentAtTailParam) {
			panic(global.GetStyledError(ExceptionAtAppendContentAtTail))
		} else if ReadFileContent(MoveContentParamLinux) == global.EmptyString {
			panic(global.GetStyledError(ExceptionAtReadFileContent))
		} else if !RenewFileCache(MoveContentParamLinux) {
			panic(global.GetStyledError(ExceptionAtRenewFileCache))
		} else if !FileContentRollBack(MoveContentParamLinux) {
			panic(global.GetStyledError(ExceptionAtFileContentRollBack))
		} else if position := SearchFileContent(MoveContentParamLinux, SearchFileContentParam); len(position) == 0 {
			panic(global.GetStyledError(ExceptionAtSearchFileContent))
		} else if !ReplaceFileContentByPosition(MoveContentParamLinux, position[0], SearchFileContentParam) {
			panic(global.GetStyledError(ExceptionAtReplaceFileContentByPosition))
		} else if !DeleteFileContentByPosition(MoveContentParamLinux, position[0]) {
			panic(global.GetStyledError(ExceptionAtDeleteFileContentByPosition))
		} else if !DeleteFileContent(MoveContentParamLinux) {
			panic(global.GetStyledError(ExceptionAtDeleteFileContent))
		} else if !RemoveFile(MoveContentParamLinux) {
			panic(global.GetStyledError(ExceptionAtRemoveFile))
		} else if !RemoveDir(CreateDirParamLinux) {
			panic(global.GetStyledError(ExceptionAtRemoveDir))
		}
	}

	// ShellOps
	// 此处不单独测试跑命令的交互环境
	if !AddCommands(AddCommandsParam) {
		panic(global.GetStyledError(ExceptionAtAddCommands))
	} else if len(GetCommands()) < 1 {
		panic(global.GetStyledError(ExceptionAtGetCommands))
	} else if AddCommands(AddCommandsParam) && !PopCommands(1) && len(GetCommands()) == 1 {
		panic(global.GetStyledError(ExceptionAtPopCommands))
	} else if !ClearCommands() {
		panic(global.GetStyledError(ExceptionAtClearCommands))
	}

	// NetOps
	if len(HttpSearch(HttpSearchParam)) < 1 {
		panic(global.GetStyledError(ExceptionAtHttpSearch))
	}

	fmt.Println(global.GetStyledSuccess(CompleteComment))
}

func TestRunTest(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("ToolTest() panic: %v", r)
		}
	}()
	ToolTest()
}
