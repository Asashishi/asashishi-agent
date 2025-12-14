package agent

const Done string = "Done"
const NeedRetry string = "Need retry"
const ToolCalls string = "tool_calls"
const ProcessingDataClean string = "Processing Data Clean 🔄"
const CallToolWithoutArgs string = "Calling %s 🔄\n"
const CallToolWithArgs string = "Calling %s 🔄 args: %s\n"
const DataCleanSysPrompt string = `
# Work
1. 对用户发来的所有信息进行摘要
2. 不要添加额外的摘要描述，对话由程序自动发送，不会要求更多信息，仅返回纯净摘要
3. 重点保留:
	- http/https 网络链接, 各类文件链接
	- 文本包含的代码片段和内容, 正确处理缩进后返回
	- 库名称, 版本和描述信息
4. 返回 ai 友好的易于理解的文本
`
const DataCleanUserPrompt string = "\n详细内容摘要, 如果内容包含代码, 保留缩进正确的代码示例"
