package start

import (
	"asashishi-agent/agent"
	"asashishi-agent/cmd"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"asashishi-agent/test"
	"asashishi-agent/tools"
	"fmt"
	"os"
	"strings"
)

func WithCliMode() {

	InitCli()

	var (
		ok          bool
		firstInput  bool
		err         error
		msg         string
		input       string
		cmdToJudge  []string
		cmdTool     func(...any) any
		isWaitInput bool              = true
		cli         agent.AgentClient = agent.AgentClient{}
	)
	cli.Init(tools.GetToolsInfo())
	if len(os.Args) > 1 && os.Args[1] == global.TestParam {
		test.RunTest()
	} else {
		for {
			select {
			case msg = <-cli.StreamChan:
				fmt.Printf(global.AIOutput, msg)
			case err = <-cli.ErrorChan:
				panic(global.GetStyledError(err.Error()))
			case input = <-global.UInput.ProcessStdin:
				go func() {
					if input != global.EmptyString {
						input = strings.TrimSpace(input)
						cmdToJudge = strings.Split(input, global.SpaceString)
						if len(cmdToJudge) > 1 && cmdToJudge[0] == global.Cmd {
							if cmdTool, ok = cmd.CmdTools[cmdToJudge[1]]; ok {
								cmdTool(input)
							}
						} else {
							fmt.Println(global.Loading)
							cli.StreamChat(input)
						}
						isWaitInput = true
					}
				}()
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
}
