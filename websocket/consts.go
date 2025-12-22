package websocket

const WebSocketPingDelay int = 1 // sec

const ClientExit string = "Client Exit"
const ProcessExit string = "Process Exit"

const WebsocketURL string = "ws://localhost:%d/ws"

// Input Type
const AIOutputType string = "ai_msg"
const SystMsgType string = "sys_msg"
const SystWarnType string = "sys_warn"
const SysErrorType string = "sys_error"
const UserInputType string = "user_input"
const ChildProcessOutputType string = "exec_output"
