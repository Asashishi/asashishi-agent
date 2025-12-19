package agent

import (
	"asashishi-agent/conf"
	"testing"

	"github.com/openai/openai-go/v3"
)

func TestToolSwitch(t *testing.T) {
	var cases = []struct {
		n string
		a string
		c *AgentClient
		e string
	}{
		{
			"GetFormatedTime",
			"",
			nil,
			"20251219122015",
		},
	}
	for _, c := range cases {
		if result := ToolCallSwitch(c.n, c.a, c.c); len(result) != len(c.e) {
			t.Errorf("%s(%s,%s,%v) = %d; expect: %d", "GetFormatedTime", c.n, c.a, c.c, len(result), len(c.e))
		}
	}
}

func TestIsUseNewLine(t *testing.T) {
	var cases = []struct {
		a *AgentClient
		e bool
	}{
		{
			&AgentClient{
				MsgContext: []openai.ChatCompletionMessageParamUnion{
					openai.SystemMessage(conf.Env.SysPrompt),
				},
			},
			false,
		},
	}
	for _, c := range cases {
		if result := IsUseNewLine(c.a); result != false {
			t.Errorf("%s(%v) = %t; expect: %t", "IsUseNewLine", c.a, result, c.e)
		}
	}
}
