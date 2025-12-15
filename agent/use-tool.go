package agent

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"fmt"

	"github.com/openai/openai-go/v3"
)

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
