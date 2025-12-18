package entry

import (
	"asashishi-agent/agent"
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"asashishi-agent/test"
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
	if len(os.Args) > 1 && os.Args[1] == global.TestParam {
		test.RunTest()
	} else {
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
					continue
				} else if err = json.Unmarshal(recved, &data); err != nil {
					conn = nil
					fmt.Println(global.GetStyledWarn(err.Error()))
					continue
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
					if conn != nil {
						if innerErr = conn.Write(
							ctx,
							ws.MessageText,
							jsonMsg,
						); innerErr != nil {
							cli.StreamForceStop = true
						}
					} else {
						cli.StreamForceStop = true
					}
				case innerErr = <-cli.ErrorChan:
					if jsonMsg, innerErr = json.Marshal(websocket.WebsocketMsg{
						Content: innerErr.Error(),
						Type:    websocket.SysErrorType,
					}); innerErr != nil {
						fmt.Println(global.GetStyledWarn(innerErr.Error()))
					}
					if conn != nil {
						if innerErr = conn.Write(
							ctx,
							ws.MessageText,
							jsonMsg,
						); innerErr != nil {
							fmt.Println(global.GetStyledWarn(innerErr.Error()))
						}
					} else {
						panic(global.GetStyledError(innerErr.Error()))
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

		if err = http.ListenAndServe(fmt.Sprintf(":%d", conf.Env.HttpPort), nil); err != nil {
			panic(global.GetStyledError(err.Error()))
		}
	}
}
