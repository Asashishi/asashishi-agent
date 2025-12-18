package websocket

type WebsocketMsg struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
