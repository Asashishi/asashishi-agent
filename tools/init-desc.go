package tools

import (
	"asashishi-agent/conf"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/shared"
)

func GetToolsInfo() []openai.ChatCompletionToolUnionParam {
	var descriptions []openai.ChatCompletionToolUnionParam = []openai.ChatCompletionToolUnionParam{

		// TimeOps
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "GetFormatedTime",
					Description: openai.String(`
					1. 此函数用于获取当前系统时间
					2. 返回一个 YYYYMMDDhhmmss 字符串
					`),
					Parameters: shared.FunctionParameters{
						"type":       "object",
						"properties": map[string]map[string]string{},
						"required":   []string{},
					},
				},
			},
		},

		// FileOps
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "GetFileList",
					Description: openai.String(`
					1. 此函数用于特定目录下的所有文件路径
					2. 如果传入 path = "" 将获取整个项目下的文件路径, 如果传入 path = 特定目录, 则获取特定目录下的文件路径
					3. 返回一个字符串数组，包含所有文件的绝对路径
					4. 第1次问答后或调用 RenameContent, CreateDir, RemoveDir, CreateFile, RemoveFile 后, 应该调用此函数
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "文件目录绝对路径",
							},
						},
						"required": []string{"path"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "RenameContent",
					Description: openai.String(`
					1. 此函数用于修改文件或文件目录的名称
					2. 函数接收两个参数, 第一个为旧文件或目录的绝对路径, 第二个为新的文件或目录的绝对路径
					4. 调用此函数后应该调用 GetFileList 获取新的文件目录信息
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"opath": {
								"type":        "string",
								"description": "要改名的旧文件或目录的绝对路径",
							},
							"npath": {
								"type":        "string",
								"description": "带新名称的绝对路径",
							},
						},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "CreateDir",
					Description: openai.String(`
					1. 此函数用于创建单层空目录
					2. 只能传入目录路径，不能包含文件名
					3. 在创建文件前，如果目录不存在，需要先调用此函数创建目录
					4. 使用此函数后, 应该调用 GetFileList 获取最新的文件目录
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要创建的目录绝对路径，不能包含文件名",
							},
						},
						"required": []string{"path"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "RemoveDir",
					Description: openai.String(`
					1. 此函数用于删除目录及其所有内容
					2. 使用此函数后, 应该调用 GetFileList 获取最新的文件目录
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要删除的目录绝对路径",
							},
						},
						"required": []string{"path"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "CreateFile",
					Description: openai.String(`
					1. 此函数用于创建空文件
					2. 确保目录存在的情况下再创建文件, 如果目录不存在, 需要先调用 CreateDir 创建目录
					3. 使用此函数后, 应该调用 GetFileList 获取最新的文件目录
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要创建的文件绝对路径，必须包含文件名",
							},
						},
						"required": []string{"path"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "RemoveFile",
					Description: openai.String(`
					1. 此函数用于删除单个文件
					3. 使用此函数后, 应该调用 GetFileList 获取最新的文件目录
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要删除的文件绝对路径",
							},
						},
						"required": []string{"path"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "ReadFileContent",
					Description: openai.String(`
					1. 此函数用于读取文件内容
					2. 文件路径来自 GetFileList 返回的列表
					3. 函数返回分按行分割的文件内容字符串数组
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要读取的文件绝对路径",
							},
						},
						"required": []string{"path"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "AppendContentAtTail",
					Description: openai.String(`
					1. 此函数用于在文件末尾追加多行内容, 写空文件时优先调用此函数
					2. 此函数接收文件的绝对路径和要写入的内容字符串作为参数
					3. 完成修改后, 使用 ReadFileContent 检查无误后 ReneFileCache, 有误则继续修改或 FileContentRollBack
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要写入的文件绝对路径",
							},
							"content": {
								"type":        "string",
								"description": `要写入的内容字符串，严格保持按 \n 换行`,
							},
						},
						"required": []string{"path", "content"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "AppendContentAtMiddle",
					Description: openai.String(`
					1. 此函数用于在文件中间追加多行内容, 在文件中间追加内容但不修改内容时优先调用此函数
					3. 此函数接收文件的绝对路径，在哪一行之后增加内容和要写入的内容作为参数
					4. 完成修改后, 使用 ReadFileContent 检查无误后 ReneFileCache, 有误则继续修改或 FileContentRollBack
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要写入的文件绝对路径",
							},
							"stp": {
								"type":        "integer",
								"description": "要插入的位置",
							},
							"content": {
								"type":        "string",
								"description": `要写入的内容字符串，严格保持按 \n 换行`,
							},
						},
						"required": []string{"path", "content"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "DeleteFileContent",
					Description: openai.String(`
					1. 此函数用于清空文件内容，但文件本身不会被删除
					2. 完成修改后, 使用 ReadFileContent 检查无误后 ReneFileCache 有误则继续修改或 FileContentRollBack
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要清空内容的文件绝对路径",
							},
						},
						"required": []string{"path"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "FileContentRollBack",
					Description: openai.String(`
					1. 此函数用于回滚文件内容到上一次缓存状态, 当其他写操作出现问题可考虑使用此函数或直接继续修改
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要回滚的文件绝对路径",
							},
						},
						"required": []string{"path"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "RenewFileCache",
					Description: openai.String(`
					1. 此函数用于更新文件内容缓存, 在写操作确认无误后必须调用
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"path": {
								"type":        "string",
								"description": "要更新缓存的文件绝对路径",
							},
						},
						"required": []string{"path"},
					},
				},
			},
		},

		// ShellOps
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "GetCommands",
					Description: openai.String(`
					1. 此函数用于获取当前命令组
					2. 此函数不接受参数
					3. 执行完成后会以字符串形式返回命令组，可以据此判断命令是否符合操作
					`),
					Parameters: shared.FunctionParameters{
						"type":       "object",
						"properties": map[string]map[string]string{},
						"required":   []string{},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "AddCommands",
					Description: openai.String(`
					1. 此函数用于向命令组添加命令
					2. 此函数接收单条命令或 "命令1 ; 命令2 ..." 的形式的命令
					3. 执行完成后可以使用 GetCommands 进行确认，可以据此判断命令是否符合操作
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"command": {
								"type":        "string",
								"description": "单条命令或 '命令1 ; 命令2 ...'' 的形式的命令",
							},
						},
						"required": []string{"command"},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "PopCommands",
					Description: openai.String(`
					1. 此函数用于按倒叙移除命令组中的的命令
					2. 此函数接收一个整数作为参数，如要移除命令组末尾的两个命令，传入 { num: 2 }
					3. 执行完成后可以使用 GetCommands 进行确认, 可以据此判断命令是否符合操作
					4. 执行 Excute 会自动执行并清空命令组, 不用连锁调用此命令
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"num": {
								"type":        "integer",
								"description": "要移除的命令个数",
							},
						},
						"required": []string{},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "ClearCommands",
					Description: openai.String(`
					1. 此函数用于清空当前命令组
					2. 此函数不接受参数
					3. 执行完成后可以使用 GetCommands 进行确认，可以据此判断命令是否符合操作
					4. 执行 Excute 会自动执行并清空命令组, 不用连锁调用此命令
					`),
					Parameters: shared.FunctionParameters{
						"type":       "object",
						"properties": map[string]map[string]string{},
						"required":   []string{},
					},
				},
			},
		},
		{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "Excute",
					Description: openai.String(`
					1. 此函数用于执行命令组，在使用之前需要通过 GetCommands 或其他命令组操作来获取具体有哪些命令组
					2. 此函数不接受参数
					3. 执行此函数后会自动清空命令组内容，无需使用其他函数清理，如有其他操作则需要再次添加命令进组后使用
					4. 执行完成后会以字符串形式返回输出执行结果，可以据此纠错或判断是否完成操作
					`),
					Parameters: shared.FunctionParameters{
						"type":       "object",
						"properties": map[string]map[string]string{},
						"required":   []string{},
					},
				},
			},
		},
	}

	if conf.Env.UseWebSearch {
		descriptions = append(descriptions, openai.ChatCompletionToolUnionParam{
			OfFunction: &openai.ChatCompletionFunctionToolParam{
				Function: shared.FunctionDefinitionParam{
					Name: "WebContentSearch",
					Description: openai.String(`
					1. 此函数用于获取相关网页的所有不含标签的文本内容
					2. 函数接收网页的 url 作为参数，返回文本字符串
					3. 通常，你可能需要根据文本内容提取符合用户输入的某些信息来连续调用这个函数
					4. 注意, 仅支持 get 请求, 如用户未指定信息源，尽量选择 库官网, github, 或其他机器人友好的网站
					`),
					Parameters: shared.FunctionParameters{
						"type": "object",
						"properties": map[string]map[string]string{
							"url": {
								"type":        "string",
								"description": "资源链接",
							},
						},
						"required": []string{"url"},
					},
				},
			},
		})
	}

	return descriptions
}
