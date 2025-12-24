package agent

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"context"
	"fmt"
	"strings"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func (cli *AgentClient) Init(
	context context.Context,
	toolList []openai.ChatCompletionToolUnionParam,
) {
	cli.Context = context
	cli.ToolsList = toolList
	cli.ModelName = conf.Env.ModelName
	cli.ErrorChan = make(chan error)
	cli.StreamChan = make(chan string)
	cli.MsgContext = make([]openai.ChatCompletionMessageParamUnion, 1)
	cli.MsgContext[0] = openai.SystemMessage(conf.Env.SysPrompt)
	cli.LlmClient = openai.NewClient(
		option.WithAPIKey(conf.Env.ApiKey),
		option.WithBaseURL(conf.Env.ModleBaseURL),
		option.WithJSONSet("temperature", conf.Env.Temperature),
		option.WithJSONSet("max_tokens", conf.Env.MaxResponseTokenLength),
	)
}

func (cli *AgentClient) ChatForWebSearchContentDataClean(prompt string) string {
	var (
		err  error
		resp *openai.ChatCompletion
	)
	if prompt == global.EmptyString {
		return ""
	} else if resp, err = cli.LlmClient.Chat.Completions.New(
		cli.Context,
		openai.ChatCompletionNewParams{
			Model: cli.ModelName,
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.SystemMessage(DataCleanSysPrompt),
				openai.UserMessage(prompt + DataCleanUserPrompt),
			},
		},
	); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
	return resp.Choices[0].Message.Content
}

func (cli *AgentClient) RemoveAllContext() {
	cli.MsgContext = cli.MsgContext[0:1]
}

func (cli *AgentClient) StreamChat(prompt string) {
	var (
		assistantMsg        string
		call                *ToolCall
		assistantMsgBuilder strings.Builder
		currTools           = map[string]*ToolCall{}
		chunk               openai.ChatCompletionChunk
	)
	if prompt != global.EmptyString {
		cli.MsgContext = append(cli.MsgContext, openai.UserMessage(prompt))
	}
	cli.CurrStrem = cli.LlmClient.Chat.Completions.NewStreaming(
		cli.Context,
		openai.ChatCompletionNewParams{
			Model:    cli.ModelName,
			Tools:    cli.ToolsList,
			Messages: cli.MsgContext,
		},
	)
	for cli.CurrStrem.Next() {
		chunk = cli.CurrStrem.Current()
		if chunk.Usage.PromptTokens > conf.Env.ContextLength*global.BitK {
			fmt.Println(global.SpaceString + global.GetStyledSystemComent(ContextOutofRange))
			cli.RemoveAllContext()
			return
		}
		if len(chunk.Choices) > 0 {
			for _, choice := range chunk.Choices {
				if choice.Delta.Content != global.EmptyString {
					cli.StreamChan <- choice.Delta.Content
					assistantMsgBuilder.WriteString(choice.Delta.Content)
				}
				if len(choice.Delta.ToolCalls) > 0 {
					for _, info := range choice.Delta.ToolCalls {
						if choice.FinishReason == ToolCalls || len(currTools) == 0 {
							call = &ToolCall{}
						}
						if info.ID != global.EmptyString {
							_, ok := currTools[info.ID]
							if !ok {
								currTools[info.ID] = call
							}
						}
						if info.Function.Name != global.EmptyString {
							call.Name = info.Function.Name
						}
						if info.Function.Arguments != global.EmptyString {
							call.Augments = call.Augments + info.Function.Arguments
						}
					}
				}
			}
		}
	}
	defer cli.CurrStrem.Close()
	if cli.CurrStrem.Err() != nil {
		cli.ErrorChan <- cli.CurrStrem.Err()
		return
	}
	assistantMsg = assistantMsgBuilder.String()
	if assistantMsg != global.EmptyString {
		cli.MsgContext = append(cli.MsgContext, openai.AssistantMessage(assistantMsg))
	}
	for k, v := range currTools {
		if k == global.EmptyString {
			continue
		}
		UseTool(k, v.Name, v.Augments, cli)
	}
	if len(currTools) > 0 {
		cli.StreamChat("")
	}
}
