package entry

import (
	"asashishi-agent/agent"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"asashishi-agent/tools"
	"asashishi-agent/websocket"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	ws "github.com/coder/websocket"
)

func WithWebMode() {
	InitWeb()
	var (
		err        error
		dirPath    string
		conn       *ws.Conn
		fileServer http.Handler
		ctx        context.Context
		mux        *http.ServeMux
		cli        agent.AgentClient = agent.AgentClient{}
	)
	mux = http.NewServeMux()
	ctx = context.Background()
	cli.Init(ctx, tools.GetToolsInfo())

	if dirPath, err = os.Getwd(); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
	fileServer = http.FileServer(http.Dir(filepath.Join(dirPath, conf.Env.ServerRootPath)))
	mux.Handle(global.HttpRootPath, fileServer)

	mux.HandleFunc(conf.Env.WebsocketRoute, func(writer http.ResponseWriter, reader *http.Request) {
		conn = websocket.GetWebsocketConn(ctx, conn, reader, writer)
	})
	defer conn.Close(ws.StatusNormalClosure, websocket.ProcessExit)

	go func() {
		var (
			err    error
			recved []byte
			data   websocket.WebsocketMsg
		)
		for {
			if conn == nil {
				global.WaitNextFrame(conf.Env.TickPerSec)
			} else if _, recved, err = conn.Read(ctx); err != nil {
				fmt.Println(global.GetStyledWarn(err.Error()))
				conn.Close(ws.StatusNormalClosure, websocket.ClientExit)
				if cli.CurrStrem != nil {
					cli.CurrStrem.Close()
				}
				conn = nil
			} else if err = json.Unmarshal(recved, &data); err != nil {
				fmt.Println(global.GetStyledWarn(err.Error()))
			} else if data.Type == websocket.UserInputType {
				global.UInput.WebsocketReadChan <- data.Content
			}
			global.WaitNextFrame(conf.Env.TickPerSec)
		}
	}()

	go func() {
		var (
			innerErr       error
			jsonMsg        []byte
			msg            string
			input          string
			processingFlag bool = false
		)
		for {
			select {
			case msg = <-cli.StreamChan:
				if jsonMsg, innerErr = json.Marshal(websocket.WebsocketMsg{
					Content: msg,
					Type:    websocket.AIOutputType,
				}); innerErr != nil {
					fmt.Println(global.GetStyledWarn(innerErr.Error()))
				} else {
					WriteOutputToWebWithRetry(ctx, conn, &cli, jsonMsg, true, true, false)
				}
			case innerErr = <-cli.ErrorChan:
				if jsonMsg, innerErr = json.Marshal(websocket.WebsocketMsg{
					Content: innerErr.Error(),
					Type:    websocket.SysErrorType,
				}); innerErr != nil {
					fmt.Println(global.GetStyledWarn(innerErr.Error()))
				} else {
					WriteOutputToWebWithRetry(ctx, conn, &cli, jsonMsg, true, false, false)
				}
			case msg = <-global.ScpOutputChan:
				if jsonMsg, innerErr = json.Marshal(websocket.WebsocketMsg{
					Content: msg,
					Type:    websocket.ChildProcessOutputType,
				}); innerErr != nil {
					fmt.Println(global.GetStyledWarn(innerErr.Error()))
				} else {
					WriteOutputToWebWithRetry(ctx, conn, &cli, jsonMsg, true, false, true)
				}
			case input = <-global.UInput.ProcessStdin:
				if !processingFlag {
					processingFlag = true
					go func() {
						cli.StreamChat(input)
						processingFlag = false
					}()
				}
			default:
				global.WaitNextFrame(conf.Env.TickPerSec)
			}
		}
	}()
	fmt.Println(
		global.GetStyledSuccess(
			fmt.Sprintf(global.WebServerStartComment, conf.Env.ServerBaseURL),
		),
	)
	go OpenBrowser(conf.Env.ServerBaseURL)
	if err = http.ListenAndServe(conf.Env.ServerListen, WithCORS(mux)); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
}
