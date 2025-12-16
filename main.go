package main

import (
	"bufio"
	"fmt"
	"strings"

	"asashishi-agent/agent"
	"asashishi-agent/backup"
	"asashishi-agent/cmd"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"asashishi-agent/test"
	"asashishi-agent/tools"
	"os"
)

func init() {
	global.SetTerminalTitle()
	conf.InitConfig()
	if conf.Env.BackUp {
		backup.BackupFiles()
	}
	global.PrintAppBanner(conf.Env.Version)
}

func main() {
	var (
		ok          bool
		firstInput  bool
		isWaitInput bool
		err         error
		msg         string
		input       string
		cmdToJudge  []string
		reader      *bufio.Reader
		cmdTool     func(...any) any
		cli         agent.AgentClient = agent.AgentClient{}
	)
	cli.Init(
		conf.Env.ApiKey,
		conf.Env.BaseURL,
		conf.Env.ModelName,
		conf.Env.SysPrompt,
		tools.GetToolsInfo(),
	)
	if len(os.Args) > 1 && os.Args[1] == global.TestParam {
		test.RunTest()
	} else {
		for {
			select {
			case msg = <-cli.StreamChan:
				fmt.Printf(global.AIOutput, msg)
			case err = <-cli.ErrorChan:
				panic(err)
			default:
				if !isWaitInput {
					isWaitInput = true
					go func() {
						if !firstInput && conf.Env.BackUp {
							fmt.Print(global.Input)
							firstInput = true
						} else {
							fmt.Print(global.InputWidthLineBreakFirst)
						}
						reader = bufio.NewReader(os.Stdin)
						if input, err = reader.ReadString(global.LineBreakChar); err != nil {
							return
						} else if input != global.EmptyString {
							input = strings.TrimSpace(input)
							cmdToJudge = strings.Split(input, global.SpaceString)
							if len(cmdToJudge) > 2 && cmdToJudge[0] == global.Cmd {
								if cmdTool, ok = cmd.CmdTools[cmdToJudge[1]]; ok {
									cmdTool(input)
								}
							} else {
								fmt.Println(global.Loading)
								cli.StreamChat(input)
							}
							isWaitInput = false
						}
					}()
				}
				global.WaitNextFrame(conf.Env.TickPerSec)
			}
		}
	}
}
