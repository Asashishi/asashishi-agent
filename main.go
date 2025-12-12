package main

import (
	"bufio"
	"fmt"
	"time"

	"asashishi-agent/agent"
	"asashishi-agent/backup"
	"asashishi-agent/conf"
	"asashishi-agent/tools"
	"os"
)

func init() {
	conf.InitConfig()
	if conf.Env.BackUP {
		backup.BackupFiles()
	}
	fmt.Printf("-- Asashishi Agent v%s --", conf.Env.Version)
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
						fmt.Print("Input: ")
						firstInput = true
					} else {
						fmt.Print("\nInput: ")
					}
					reader = bufio.NewReader(os.Stdin)
					if input, err = reader.ReadString('\n'); err != nil {
						return
					} else if input != "" {
						fmt.Print("Loading...\n")
						cli.StreamChat(input)
						isWaitInput = false
					}
				}()
			} else {
				time.Sleep(
					(time.Duration((1000 / conf.Env.TickPerSec) * 1000)) * time.Microsecond,
				)
			}
		}
	}
}
