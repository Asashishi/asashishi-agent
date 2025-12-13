package main

import (
	"bufio"
	"fmt"
	"time"

	"asashishi-agent/agent"
	"asashishi-agent/backup"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"asashishi-agent/test"
	"asashishi-agent/tools"
	"os"
)

func wait() {
	time.Sleep(
		(time.Duration((1000 / conf.Env.TickPerSec) * global.FloatK)) * time.Microsecond,
	)
}

func init() {
	conf.InitConfig()
	if conf.Env.BackUP {
		backup.BackupFiles()
	}
	fmt.Printf(global.AsashishiAgentWithVersion, conf.Env.Version)
}

func main() {
	var (
		firstInput  bool
		isWaitInput bool
		err         error
		msg         string
		input       string
		reader      *bufio.Reader
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
				fmt.Print(msg)
			case err = <-cli.ErrorChan:
				panic(err)
			default:
				if !isWaitInput {
					isWaitInput = true
					go func() {
						if !firstInput && conf.Env.BackUP {
							fmt.Print(global.Input)
							firstInput = true
						} else {
							fmt.Print(global.InputWidthLineBreakFirst)
						}
						reader = bufio.NewReader(os.Stdin)
						if input, err = reader.ReadString(global.LineBreakChar); err != nil {
							return
						} else if input != global.EmptyString {
							fmt.Println(global.Loading)
							cli.StreamChat(input)
							isWaitInput = false
						}
					}()
				}
				wait()
			}
		}
	}
}
