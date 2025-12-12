package agent

import (
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
		fmt.Println("Processing Data Clean", "🔄")
		if searchResult == "" || searchResult == NotFound {
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
	case "RenameContent":
		var args struct {
			OPath string `json:"opath"`
			NPath string `json:"nPath"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if tools.RenameContent(args.OPath, args.NPath) {
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
		if data, err := json.Marshal(tools.ReadFileContent(args.Path)); err != nil {
			message = NeedRetry
		} else {
			message = string(data)
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
	case "AppendContentAtMiddle":
		var args struct {
			Path    string `json:"path"`
			Stp     int    `json:"stp"`
			Content string `json:"content"`
		}
		json.Unmarshal([]byte(arguments), &args)
		if data, err := json.Marshal(tools.AppendContentAtMiddle(args.Path, args.Stp, args.Content)); err != nil {
			message = NeedRetry
		} else {
			message = string(data)
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
	case "Excute":
		message = tools.Execute()
	}
	return message
}
