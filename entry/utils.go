package entry

import (
	"asashishi-agent/agent"
	"asashishi-agent/backup"
	"asashishi-agent/cmd"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"asashishi-agent/websocket"
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	ws "github.com/coder/websocket"
)

func InitCli() {
	global.SetTerminalTitle()
	if conf.Env.BackUp {
		backup.BackupFiles(conf.Env.System)
	}
	global.PrintAppBanner(conf.Env.Version)
	go global.InitGlobalCliUInput()
}

func InitWeb() {
	global.SetTerminalTitle()
	if conf.Env.BackUp {
		backup.BackupFiles(conf.Env.System)
	}
	global.PrintAppBanner(conf.Env.Version)
	go global.InitGlobalWebUInput()
}

func OpenBrowser(url string) {
	if conf.Env.System == conf.Windows {
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Run()
	} else {
		exec.Command("xdg-open", url).Run()
	}
}

func WithCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, reader *http.Request) {
		// 设置允许所有跨域
		writer.Header().Set(ServerAllowOrigin, conf.Env.AllowOrigin)
		writer.Header().Set(ServerAllowHeaders, conf.Env.AllowHeaders)
		writer.Header().Set(ServerAllowMethods, conf.Env.AllowMethods)
		if reader.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(writer, reader)
	})
}

func HandleCliUInput(input string, cli *agent.AgentClient, isWaitInput *bool) {
	var (
		ok         bool
		cmdToJudge []string
		cmdTool    func(...string)
	)
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
		*isWaitInput = true
	}
}

func WriteOutputToWebWithRetry(ctx context.Context, conn *ws.Conn, cli *agent.AgentClient, jsonMsg []byte, tConnC bool, tStreamC bool) {
	var (
		err      error
		errCount int = 0
	)
	if conn != nil {
		for {
			if err = conn.Write(
				ctx,
				ws.MessageText,
				jsonMsg,
			); err != nil {
				if errCount == 3 {
					if tConnC {
						conn.Close(ws.StatusNormalClosure, websocket.ClientExit)
					}
					fmt.Println(global.GetStyledWarn(err.Error()))
				}
				errCount++
				time.Sleep(time.Millisecond * time.Duration(WebsocketWriteRetryDelay))
			} else {
				break
			}
		}
	} else if tStreamC {
		cli.CurrStrem.Close()
	}

}
