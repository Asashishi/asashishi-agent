package start

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
		msg        string
		dirPath    string
		conn       *ws.Conn
		fileServer http.Handler
		ctx        context.Context
		cli        agent.AgentClient = agent.AgentClient{}
	)

	cli.Init(tools.GetToolsInfo())

	if dirPath, err = os.Getwd(); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
	fileServer = http.FileServer(http.Dir(filepath.Join(dirPath, conf.Env.ServerRootPath)))
	http.Handle("/", fileServer)

	fmt.Println(
		global.GetStyledSuccess(
			fmt.Sprintf(global.WebServerStartComment, conf.Env.HttpPort),
		),
	)
	ctx, conn = websocket.WebSocketServerInit()
	defer conn.Close(ws.StatusNormalClosure, websocket.ProcessExit)

	go func() {
		var (
			err    error
			recved []byte
			data   websocket.WebsocketMsg
		)
		for {
			if _, recved, err = conn.Read(ctx); err != nil {
				fmt.Println(global.GetStyledWarn(err.Error()))
				continue
			} else if err = json.Unmarshal(recved, &data); err != nil {
				fmt.Println(global.GetStyledWarn(err.Error()))
				continue
			} else if data.Type == websocket.UserInputType {
				global.UInput.WebsocketReadChan <- data.Content
			}
		}
	}()

	go func() {
		var (
			innerErr error
			jsonMsg  []byte
		)
		for {
			select {
			case msg = <-cli.StreamChan:
				if jsonMsg, innerErr = json.Marshal(websocket.WebsocketMsg{
					Content: msg,
					Type:    websocket.AIOutputType,
				}); err != nil {
					fmt.Println(global.GetStyledWarn(innerErr.Error()))
				}
				conn.Write(
					ctx,
					ws.MessageText,
					jsonMsg,
				)
			case err = <-cli.ErrorChan:
				if jsonMsg, innerErr = json.Marshal(websocket.WebsocketMsg{
					Content: err.Error(),
					Type:    websocket.SysErrorType,
				}); innerErr != nil {
					panic(global.GetStyledError(innerErr.Error()))
				}
				conn.Write(
					ctx,
					ws.MessageText,
					jsonMsg,
				)
			default:
				global.WaitNextFrame(conf.Env.TickPerSec)
			}
		}
	}()

	if err = http.ListenAndServe(fmt.Sprintf(":%d", conf.Env.HttpPort), nil); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
}
