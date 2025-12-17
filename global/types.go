package global

type GlobalUInput struct {
	IsChildProcess    bool
	ProcessStdin      chan string
	ChildProcessStdin chan string
}
