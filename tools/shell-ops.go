package tools

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
)

var Commands []string = []string{}

func killChildProcessGroup(pid int) {
	if conf.Env.System == conf.Windows {
		exec.Command("taskkill", "/PID", strconv.Itoa(pid), "/F", "/T").Run()
	} else {
		exec.Command("kill", "-9", fmt.Sprintf("-%d", pid)).Run()
	}
}

func buildShell() *exec.Cmd {
	var (
		shell           *exec.Cmd
		fomatedCommands string
	)
	if conf.Env.System == conf.Windows {
		fomatedCommands = WindowsInitialCommand
		for _, cmd := range Commands {
			fomatedCommands += fmt.Sprintf("; if ($?) { %s }", cmd)
		}
		shell = exec.Command("powershell", "-Command", fomatedCommands)
	} else {
		fomatedCommands = LinuxInitialCommand
		for _, cmd := range Commands {
			fomatedCommands += fmt.Sprintf(" && %s", cmd)
		}
		shell = exec.Command("bash", "-c", fomatedCommands)
	}
	Commands = []string{}
	return shell
}

func GetCommands() []string {
	return Commands
}

func AddCommands(command string) bool {
	Commands = append(Commands, command)
	return true
}

func PopCommands(num int) bool {
	Commands = Commands[0 : len(Commands)-num]
	return true
}

func ClearCommands() bool {
	Commands = []string{}
	return true
}

func InterActiveExecute() string {
	var (
		err       error
		shell     *exec.Cmd
		buffer    bytes.Buffer
		stdinPipe io.WriteCloser
	)
	shell = buildShell()
	if stdinPipe, err = shell.StdinPipe(); err != nil {
		return err.Error()
	}
	global.UInput.IsChildProcess = true
	shell.Stdout = io.MultiWriter(os.Stdout, &buffer)
	shell.Stderr = io.MultiWriter(os.Stderr, &buffer)
	go func() {
		var msg string
		for {
			select {
			case msg = <-global.UInput.ChildProcessStdin:
				stdinPipe.Write([]byte(msg))
			default:
				if !global.UInput.IsChildProcess {
					break
				}
				global.WaitNextFrame(conf.Env.TickPerSec)
			}
		}
	}()
	err = shell.Run()
	global.UInput.IsChildProcess = false
	if err != nil {
		return err.Error()
	}
	return buffer.String()
}

func NoInterActiveExecute() string {
	defer fmt.Println(global.GetStyledSystemComent(Exit))
	fmt.Println(global.GetStyledSystemComent(PressEnterToBackToChat))
	var (
		stopFlag bool
		err      error
		shell    *exec.Cmd
		buffer   bytes.Buffer
	)
	stopFlag = false
	shell = buildShell()
	global.UInput.IsChildProcess = true
	shell.Stdout = io.MultiWriter(os.Stdout, &buffer)
	shell.Stderr = io.MultiWriter(os.Stderr, &buffer)
	go func() {
		var msg string
		for {
			select {
			case msg = <-global.UInput.ChildProcessStdin:
				if msg != "" {
					stopFlag = true
					killChildProcessGroup(shell.Process.Pid)
				}
			default:
				if !global.UInput.IsChildProcess {
					break
				}
				global.WaitNextFrame(conf.Env.TickPerSec)
			}
		}
	}()
	err = shell.Run()
	global.UInput.IsChildProcess = false
	if err != nil {
		if stopFlag {
			return buffer.String() + StopedByUser
		}
		return err.Error()
	}
	return buffer.String()
}
