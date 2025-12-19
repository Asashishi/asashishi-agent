package entry

import (
	"asashishi-agent/agent"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"asashishi-agent/tools"
	"context"
	"fmt"
)

func WithCliMode() {
	InitCli()
	var (
		firstInput  bool
		err         error
		msg         string
		input       string
		isWaitInput bool              = true
		cli         agent.AgentClient = agent.AgentClient{}
	)
	cli.Init(context.Background(), tools.GetToolsInfo())
	for {
		select {
		case msg = <-cli.StreamChan:
			fmt.Printf(global.AIOutput, msg)
		case err = <-cli.ErrorChan:
			panic(global.GetStyledError(err.Error()))
		case input = <-global.UInput.ProcessStdin:
			go HandleCliUInput(input, &cli, &isWaitInput)
		default:
			if isWaitInput {
				if !firstInput && conf.Env.BackUp {
					fmt.Print(global.Input)
					firstInput = true
				} else {
					fmt.Print(global.InputWidthLineBreakFirst)
				}
				isWaitInput = false
			}
			global.WaitNextFrame(conf.Env.TickPerSec)
		}
	}
}
