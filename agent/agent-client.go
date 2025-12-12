package agent

import (
	"asashishi-agent/conf"
	"context"
	"strings"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/packages/ssestream"
)

func (cli *AgentClient) Init(
	apiKey string,
	baseUrl string,
	modelName string,
	sysPrompt string,
	toolList []openai.ChatCompletionToolUnionParam,
) {
	cli.ToolsList = toolList
	cli.ModelName = modelName
	cli.ErrorChan = make(chan error)
	cli.StreamChan = make(chan string)
	cli.Context = context.Background()
	cli.MsgContext = make([]openai.ChatCompletionMessageParamUnion, 1)
	cli.MsgContext[0] = openai.SystemMessage(sysPrompt)
	cli.LlmClient = openai.NewClient(
		option.WithAPIKey(apiKey),
		option.WithBaseURL(baseUrl),
		option.WithJSONSet("temperature", conf.Env.Temperature),
		option.WithJSONSet("max_tokens", conf.Env.MaxResponseTokenLength),
	)
}

func (cli *AgentClient) SetContextToLastUserPrompt() {
	for i := len(cli.MsgContext) - 1; i >= 0; i-- {
		if cli.MsgContext[i].OfUser != nil {
			cli.MsgContext = cli.MsgContext[i:]
			return
		}
	}
}

func (cli *AgentClient) ChatForWebSearchContentDataClean(prompt string) string {
	var (
		err  error
		resp *openai.ChatCompletion
	)
	if prompt == "" {
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
		return ""
	}
	return resp.Choices[0].Message.Content
}

func (cli *AgentClient) StreamChat(prompt string) {
	var (
		assistantMsg        string
		call                *ToolCall
		assistantMsgBuilder strings.Builder
		currTools           = map[string]*ToolCall{}
		chunk               openai.ChatCompletionChunk
		stream              *ssestream.Stream[openai.ChatCompletionChunk]
	)
	if prompt != "" {
		cli.MsgContext = append(cli.MsgContext, openai.UserMessage(prompt))
	}
	stream = cli.LlmClient.Chat.Completions.NewStreaming(
		cli.Context,
		openai.ChatCompletionNewParams{
			Model:    cli.ModelName,
			Tools:    cli.ToolsList,
			Messages: cli.MsgContext,
		},
	)
	for stream.Next() {
		chunk = stream.Current()
		if chunk.Usage.PromptTokens > conf.Env.ContextLength*K {
			cli.SetContextToLastUserPrompt()
		}
		if len(chunk.Choices) > 0 {
			for _, choice := range chunk.Choices {
				if choice.Delta.Content != "" {
					cli.StreamChan <- choice.Delta.Content
					assistantMsgBuilder.WriteString(choice.Delta.Content)
				}
				if len(choice.Delta.ToolCalls) > 0 {
					for _, info := range choice.Delta.ToolCalls {
						if choice.FinishReason == "tool_calls" || len(currTools) == 0 {
							call = &ToolCall{
								Name:     "",
								Augments: "",
							}
						}
						if info.ID != "" {
							_, ok := currTools[info.ID]
							if !ok {
								currTools[info.ID] = call
							}
						}
						if info.Function.Name != "" {
							call.Name = info.Function.Name
						}
						if info.Function.Arguments != "" {
							call.Augments = call.Augments + info.Function.Arguments
						}
					}
				}
			}
		}
	}
	stream.Close()
	assistantMsg = assistantMsgBuilder.String()
	if assistantMsg != "" {
		cli.MsgContext = append(cli.MsgContext, openai.AssistantMessage(assistantMsg))
	}
	for k, v := range currTools {
		if k == "" {
			continue
		}
		UseTool(k, v.Name, v.Augments, cli)
	}
	if stream.Err() != nil {
		cli.ErrorChan <- stream.Err()
		return
	}
	if len(currTools) > 0 {
		cli.StreamChat("")
	}
}
