package tools

import (
	"asashishi-agent/conf"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/shared"
)

var TimeOps []openai.ChatCompletionToolUnionParam = []openai.ChatCompletionToolUnionParam{
	{
		OfFunction: &openai.ChatCompletionFunctionToolParam{
			Function: shared.FunctionDefinitionParam{
				Name: "GetFormatedTime",
				Description: openai.String(`
					1. 此函数用于获取当前系统时间
					2. 返回一个 YYYYMMDDhhmmss 字符串
					3. 仅在需要准确时间和记录日志时调用
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

var FileOps []openai.ChatCompletionToolUnionParam = []openai.ChatCompletionToolUnionParam{
	{
		OfFunction: &openai.ChatCompletionFunctionToolParam{
			Function: shared.FunctionDefinitionParam{
				Name: "GetFileList",
				Description: openai.String(`
				1. 此函数用于特定目录下的所有文件路径
				2. 返回一个字符串数组，包含所有文件的绝对路径
				3. 第1次问答后或调用 MoveContent, CreateDir, RemoveDir, CreateFile, RemoveFile 后, 应该调用此函数
				`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]string{
						"path": {
							"type":        "string",
							"description": "文件目录绝对路径, 如果传入 path = '' 将获取整个项目下的文件路径, 如果传入 path = 特定目录, 则获取特定目录下的文件路径",
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
				Name: "MoveContent",
				Description: openai.String(`
				1. 此函数用于修改文件或文件目录的名称或其位置
				2. 调用此函数后应该调用 GetFileList 获取新的文件目录信息
				`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]string{
						"opath": {
							"type":        "string",
							"description": "要改名的旧文件或目录的绝对路径, 来自 GetFileList 返回的列表",
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
							"description": "要创建的目录绝对路径",
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
							"description": "要删除的目录绝对路径, 来自 GetFileList 返回的列表",
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
				2. 使用此函数后, 应该调用 GetFileList 获取最新的文件目录
				`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]string{
						"path": {
							"type":        "string",
							"description": "要删除的文件绝对路径, 来自 GetFileList 返回的列表",
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
				2. 函数会返回文件内容的文本字符串
				`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]string{
						"path": {
							"type":        "string",
							"description": "要读取的文件绝对路径, 来自 GetFileList 返回的列表",
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
				2. 完成修改后, 使用 ReadFileContent 检查无误后 ReneFileCache, 有误则继续修改或 FileContentRollBack
				`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]string{
						"path": {
							"type":        "string",
							"description": "要写入的文件绝对路径, 来自 GetFileList 返回的列表",
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
				Name: "SearchFileContent",
				Description: openai.String(`
				1. 此函数用于查找一段内容在文件中出现的字符起始位置和结束位置
				2. 传入的 content 参数必须是整段文件内容, 而不是断续的残句或标题
				2. 内容可能出现多次, 函数会以数组的形式 如 [[9, 120], [309, 420]...] 的形式返回
				3. 得到正确的位置后, 需要将位置传递给 ReplaceFileContentByPosition 或 DeleteFileContentByPosition 进行替换或删除操作
				4. 如果已经读取过文件内容, 可以在上下文中摘选要修改的片段进行再次查找, 未有修改时不得频繁获取整个文件内容 
				`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]string{
						"path": {
							"type":        "string",
							"description": "文件绝对路径, 来自 GetFileList 返回的列表",
						},
						"content": {
							"type":        "string",
							"description": "要查找的内容字符串, 注意需要与文件内容的片段完全一致",
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
				Name: "ReplaceFileContentByPosition",
				Description: openai.String(`
				1. 此函数用于根据起始字符位置和结束字符位置来替换整段文件内容
				2. 为减少纠错次数, 调用此函数前你必须先调用 SearchFileContent
				4. 完成修改后, 使用 ReadFileContent 检查无误后 ReneFileCache 有误则继续修改或 FileContentRollBack
				`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]any{
						"path": {
							"type":        "string",
							"description": "文件绝对路径, 来自 GetFileList 返回的列表",
						},
						"position": {
							"type": "array",
							"items": map[string]string{
								"type": "integer",
							},
							"description": "长度为 2 的整数数组，格式为 [start, end]，表示文本片段的起止位置(左闭右开), 内容来自 SearchFileContent",
						},
						"content": {
							"type":        "string",
							"description": "新的, 要替换为的字符串内容内容",
						},
					},
					"required": []string{"path", "position", "content"},
				},
			},
		},
	},
	{
		OfFunction: &openai.ChatCompletionFunctionToolParam{
			Function: shared.FunctionDefinitionParam{
				Name: "DeleteFileContentByPosition",
				Description: openai.String(`
				1. 此函数用于根据起始字符位置和结束字符位置来删除整段文件内容
				2. 为减少纠错次数, 调用此函数前你必须先调用 SearchFileContent
				3. 完成修改后, 使用 ReadFileContent 检查无误后 ReneFileCache 有误则继续修改或 FileContentRollBack
				`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]any{
						"path": {
							"type":        "string",
							"description": "文件绝对路径, 来自 GetFileList 返回的列表",
						},
						"position": {
							"type": "array",
							"items": map[string]string{
								"type": "integer",
							},
							"description": "长度为 2 的整数数组，格式为 [start, end]，表示文本片段的起止位置(左闭右开), 内容来自 SearchFileContent",
						},
					},
					"required": []string{"path", "position"},
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
							"description": "要回滚的文件绝对路径, 来自 GetFileList 返回的列表",
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
							"description": "要更新缓存的文件绝对路径, 来自 GetFileList 返回的列表",
						},
					},
					"required": []string{"path"},
				},
			},
		},
	},
}

var ShellOps []openai.ChatCompletionToolUnionParam = []openai.ChatCompletionToolUnionParam{
	{
		OfFunction: &openai.ChatCompletionFunctionToolParam{
			Function: shared.FunctionDefinitionParam{
				Name: "GetCommands",
				Description: openai.String(`
					1. 此函数用于获取当前命令组
					2. 执行完成后会以字符串形式返回命令组，可以据此判断命令是否符合操作
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
					2. 所有命令添加完成后可以使用 GetCommands 进行确认，可以据此判断命令是否符合操作
					`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]string{
						"command": {
							"type":        "string",
							"description": "单条命令或 '命令1 ;(Linux环境则使用 &&) 命令2 ...'' 的形式的命令",
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
					2. 执行完成后可以使用 GetCommands 进行确认, 可以据此判断命令是否符合操作
					3. 执行命令组会清空命令组, ClearCommands 和执行命令组后不得调用此函数
					`),
				Parameters: shared.FunctionParameters{
					"type": "object",
					"properties": map[string]map[string]string{
						"num": {
							"type":        "integer",
							"description": "要移除的命令个数, 如要移除命令组末尾的两个命令，传入 { num: 2 }",
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
					4. 执行命令组会清空命令组, 执行命令组后不得调用此函数
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
				Name: "InteractiveExecute",
				Description: openai.String(`
					1. 此函数用于对交互式程序(如需要用户执行输入进行操作的程序)执行命令组，在使用之前需要通过 GetCommands 或其他命令组操作来获取具体有哪些命令组
					2. 执行此函数后会自动清空命令组内容，无需使用其他函数清理，如有其他操作则需要再次添加命令进组后使用
					3. 执行完成后会以字符串形式返回输出执行结果，可以据此纠错或判断是否完成操作
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
				Name: "NoInteractiveExecute",
				Description: openai.String(`
					1. 此函数用于对非交互式程序(如 Web 服务)执行命令组，在使用之前需要通过 GetCommands 或其他命令组操作来获取具体有哪些命令组
					2. 执行此函数后会自动清空命令组内容，无需使用其他函数清理，如有其他操作则需要再次添加命令进组后使用
					3. 执行完成后会以字符串形式返回输出执行结果，可以据此纠错或判断是否完成操作
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

var NetOps []openai.ChatCompletionToolUnionParam = []openai.ChatCompletionToolUnionParam{
	{
		OfFunction: &openai.ChatCompletionFunctionToolParam{
			Function: shared.FunctionDefinitionParam{
				Name: "HttpSearch",
				Description: openai.String(`
					1. 此函数用于获取相关网页的所有不含标签的文本内容
					2. 通常，你可能需要根据文本内容提取符合用户输入的某些信息来连续调用这个函数
					3. 注意, 仅支持 get 请求, 如用户未指定信息源，尽量选择 库官网, github, 或其他机器人友好的网站
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
	},
}

func GetToolsInfo() []openai.ChatCompletionToolUnionParam {
	var descriptions []openai.ChatCompletionToolUnionParam

	descriptions = append(descriptions, TimeOps...)
	descriptions = append(descriptions, FileOps...)
	descriptions = append(descriptions, ShellOps...)
	descriptions = append(descriptions, NetOps...)

	if conf.Env.UseWebSearch {
		descriptions = append(descriptions, NetOps...)
	}
	return descriptions
}
