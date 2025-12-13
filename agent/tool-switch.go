package agent

import (
	"asashishi-agent/global"
	"asashishi-agent/tools"
	"encoding/json"
	"fmt"
)

func ToolCallSwitch(name string, arguments string, cli *AgentClient) string {
	var message string
	switch name {
	// timeOps
	case "GetFormatedTime":
		message = tools.GetFormatedTime()

	// netOps
	case "WebContentSearch":
		var (
			args struct {
				ForceNewSearch bool   `json:"force"`
				Url            string `json:"url"`
				Intent         string `json:"intent"`
			}
			searchResult string
		)
		json.Unmarshal([]byte(arguments), &args)
		searchResult = tools.WebContentSearch(args.Url)
		fmt.Println(ProcessingDataClean)
		if searchResult == global.EmptyString {
			message = searchResult
		} else {
			message = cli.ChatForWebSearchContentDataClean(searchResult)
		}

	// fileOps
	case "GetFileList":
		var args struct {
			Path string `json:"path"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if data, err := json.Marshal(tools.GetFileList(args.Path)); err != nil {
			message = NeedRetry
		} else {
			message = string(data)
		}
	case "MoveContent":
		var args struct {
			OPath string `json:"opath"`
			NPath string `json:"nPath"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.MoveContent(args.OPath, args.NPath) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "CreateDir":
		var args struct {
			Path string `json:"path"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.CreateDir(args.Path) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "RemoveDir":
		var args struct {
			Path string `json:"path"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.RemoveDir(args.Path) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "CreateFile":
		var args struct {
			Path string `json:"path"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.CreateFile(args.Path) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "RemoveFile":
		var args struct {
			Path string `json:"path"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.RemoveFile(args.Path) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "ReadFileContent":
		var args struct {
			Path string `json:"path"`
		}
		json.Unmarshal([]byte(arguments), &args)
		var result string = tools.ReadFileContent(args.Path)
		if result != global.EmptyString {
			message = result
		} else {
			message = NeedRetry
		}
	case "AppendContentAtTail":
		var args struct {
			Path    string `json:"path"`
			Content string `json:"content"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.AppendContentAtTail(args.Path, args.Content) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "SearchFileContent":
		var args struct {
			Path    string `json:"path"`
			Content string `json:"content"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if data, err := json.Marshal(tools.SearchFileContent(args.Path, args.Content)); err != nil {
			message = NeedRetry
		} else {
			message = string(data)
		}
	case "ReplaceFileContentByPosition":
		var args struct {
			Position []int  `json:"position"`
			Path     string `json:"path"`
			Content  string `json:"content"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.ReplaceFileContentByPosition(args.Path, args.Position, args.Content) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "DeleteFileContentByPosition":
		var args struct {
			Position []int  `json:"position"`
			Path     string `json:"path"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.DeleteFileContentByPosition(args.Path, args.Position) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "DeleteFileContent":
		var args struct {
			Path string
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.DeleteFileContent(args.Path) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "RenewFileCache":
		var args struct {
			Path string `json:"path"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.RenewFileCache(args.Path) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "FileContentRollBack":
		var args struct {
			Path string `json:"path"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.FileContentRollBack(args.Path) {
			message = Done
		} else {
			message = NeedRetry
		}

	// shellOps
	case "GetCommands":
		if data, err := json.Marshal(tools.GetCommands()); err != nil {
			message = NeedRetry
		} else {
			message = string(data)
		}
	case "AddCommands":
		var args struct {
			Command string `json:"command"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.AddCommands(args.Command) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "PopCommands":
		var args struct {
			Num int `json:"num"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.PopCommands(args.Num) {
			message = Done
		} else {
			message = NeedRetry
		}
	case "ClearCommands":
		if tools.ClearCommands() {
			message = Done
		} else {
			message = NeedRetry
		}
	case "InterActiveExcute":
		message = tools.InterActiveExecute()
	case "NoInterActiveExecute":
		message = tools.NoInterActiveExecute()
	}
	return message
}
