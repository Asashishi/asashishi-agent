package agent

import (
	"context"

	"github.com/openai/openai-go/v3"
)

type ToolCall struct {
	Name     string
	Augments string
}

type AgentClient struct {
	ModelName  string
	ErrorChan  chan error
	StreamChan chan string
	LlmClient  openai.Client
	Context    context.Context
	ToolsList  []openai.ChatCompletionToolUnionParam
	MsgContext []openai.ChatCompletionMessageParamUnion
}
