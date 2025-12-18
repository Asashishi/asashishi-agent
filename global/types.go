package global

type GlobalUInput struct {
	IsChildProcess    bool
	WebsocketReadChan chan string
	ProcessStdin      chan string
	ChildProcessStdin chan string
}
