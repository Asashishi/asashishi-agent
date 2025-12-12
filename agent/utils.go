package agent

import "github.com/openai/openai-go/v3"

func IsUseNewLine(cli *AgentClient) bool {
	var item openai.ChatCompletionMessageParamUnion = cli.MsgContext[len(cli.MsgContext)-1]
	var prevItem openai.ChatCompletionMessageParamUnion = cli.MsgContext[len(cli.MsgContext)-3]
	return (item.OfTool == nil && prevItem.OfTool != nil) || item.OfAssistant != nil || item.OfUser != nil
}
