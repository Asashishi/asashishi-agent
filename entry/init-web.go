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
		cli        agent.AgentClient = agent.AgentClient{}
	)

	ctx = context.Background()
	cli.Init(ctx, tools.GetToolsInfo())

	if dirPath, err = os.Getwd(); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
	fileServer = http.FileServer(http.Dir(filepath.Join(dirPath, conf.Env.ServerRootPath)))
	http.Handle(global.HttpRootPath, fileServer)

	http.HandleFunc(conf.Env.WebsocketRoute, func(writer http.ResponseWriter, reader *http.Request) {
		conn = websocket.GetWebsocketConn(conn, writer, reader)
	})
	defer conn.Close(ws.StatusInternalError, websocket.ProcessExit)

	fmt.Println(
		global.GetStyledSuccess(
			fmt.Sprintf(global.WebServerStartComment, conf.Env.HttpPort),
		),
	)

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
				conn = nil
				cli.StreamForceStop = true
				fmt.Println(global.GetStyledWarn(err.Error()))
			} else if err = json.Unmarshal(recved, &data); err != nil {
				fmt.Println(global.GetStyledWarn(err.Error()))
			} else if data.Type == websocket.UserInputType {
				global.UInput.WebsocketReadChan <- data.Content
			}
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
				}
				WriteAIRespToWebsocketOutput(ctx, conn, &cli, jsonMsg)
			case msg = <-global.ScpOutputChan:
				if jsonMsg, innerErr = json.Marshal(websocket.WebsocketMsg{
					Content: msg,
					Type:    websocket.ChildProcessOutputType,
				}); innerErr != nil {
					fmt.Println(global.GetStyledWarn(innerErr.Error()))
				}
				WriteScpOutputToWebsocketOutput(ctx, conn, &cli, jsonMsg)
			case innerErr = <-cli.ErrorChan:
				if jsonMsg, innerErr = json.Marshal(websocket.WebsocketMsg{
					Content: innerErr.Error(),
					Type:    websocket.SysErrorType,
				}); innerErr != nil {
					fmt.Println(global.GetStyledWarn(innerErr.Error()))
				}
				WriteAIErrorToWebsocketOutput(ctx, conn, &cli, jsonMsg, innerErr)
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

	if err = http.ListenAndServe(fmt.Sprintf(":%d", conf.Env.HttpPort), nil); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
}
