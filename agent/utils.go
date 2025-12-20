package agent

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"fmt"
	"math"

	"github.com/openai/openai-go/v3"
)

func IsUseNewLine(cli *AgentClient) bool {
	var item openai.ChatCompletionMessageParamUnion = cli.MsgContext[len(cli.MsgContext)-1]
	var prevItem openai.ChatCompletionMessageParamUnion = cli.MsgContext[int(math.Max(float64(len(cli.MsgContext)-3), 0))]
	return (item.OfTool == nil && prevItem.OfTool != nil) || item.OfAssistant != nil || item.OfUser != nil
}

func UseTool(id string, name string, arguments string, cli *AgentClient) {
	var (
		toolMessage  string
		assistantMsg openai.ChatCompletionMessageParamUnion
	)
	if IsUseNewLine(cli) {
		fmt.Print(global.LineBreakString)
	}
	if conf.Env.ShowToolCallArgs {
		fmt.Printf(CallToolWithArgs, name, arguments)
	} else {
		fmt.Printf(CallToolWithoutArgs, name)
	}
	toolMessage = ToolCallSwitch(name, arguments, cli)
	assistantMsg = openai.ChatCompletionMessageParamUnion{
		OfAssistant: &openai.ChatCompletionAssistantMessageParam{
			Role: "assistant",
			ToolCalls: []openai.ChatCompletionMessageToolCallUnionParam{
				{
					OfFunction: &openai.ChatCompletionMessageFunctionToolCallParam{
						Type: "function",
						ID:   id,
						Function: openai.ChatCompletionMessageFunctionToolCallFunctionParam{
							Name:      name,
							Arguments: arguments,
						},
					},
				},
			},
		},
	}
	cli.MsgContext = append(
		cli.MsgContext,
		assistantMsg,
		openai.ToolMessage(toolMessage, id),
	)
}
