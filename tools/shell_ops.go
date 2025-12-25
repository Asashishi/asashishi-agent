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

func stopHook(shell *exec.Cmd) {
	for {
		if global.UInput.IsChildProcess == false {
			killChildProcessGroup(shell.Process.Pid)
			break
		}
		global.WaitNextFrame(conf.Env.TickPerSec)
	}
}

func buildShell() *exec.Cmd {
	var (
		fomatedCommands string
		shell           *exec.Cmd
	)
	if conf.Env.System == conf.Windows {
		fomatedCommands = WindowsInitialCommand
		for _, cmd := range global.Commands {
			fomatedCommands += fmt.Sprintf("; if ($?) { %s }", cmd)
		}
		shell = exec.Command("powershell", "-Command", fomatedCommands)
	} else {
		fomatedCommands = LinuxInitialCommand
		for _, cmd := range global.Commands {
			fomatedCommands += fmt.Sprintf(" && %s", cmd)
		}
		shell = exec.Command("bash", "-c", fomatedCommands)
	}
	global.Commands = []string{}
	return shell
}

func killChildProcessGroup(pid int) {
	if conf.Env.System == conf.Windows {
		exec.Command("taskkill", "/PID", strconv.Itoa(pid), "/F", "/T").Run()
	} else {
		exec.Command("kill", "-9", fmt.Sprintf("-%d", pid)).Run()
	}
}

func handleScpExit(flag *bool, shell *exec.Cmd) {
	for {
		select {
		case <-global.UInput.ChildProcessStdin:
			*flag = true
			killChildProcessGroup(shell.Process.Pid)
		default:
			if !global.UInput.IsChildProcess {
				break
			}
			global.WaitNextFrame(conf.Env.TickPerSec)
		}
	}
}

func handleScpInteractiveInput(stdinPipe *io.WriteCloser) {
	var msg string
	for {
		select {
		case msg = <-global.UInput.ChildProcessStdin:
			if conf.Env.WebMode {
				(*stdinPipe).Write([]byte(msg + "\n"))
			} else {
				(*stdinPipe).Write([]byte(msg))
			}
		default:
			if !global.UInput.IsChildProcess {
				break
			}
			global.WaitNextFrame(conf.Env.TickPerSec)
		}
	}
}

func handleScpOutputForWeb(reader *io.PipeReader) {
	var (
		length     int
		innerErr   error
		buffer     []byte
		errContent string
	)
	for {
		buffer = make([]byte, 4096)
		length, innerErr = reader.Read(buffer)
		global.ScpOutputChan <- string(buffer[:length])
		if innerErr != nil {
			errContent = innerErr.Error()
			if innerErr != io.EOF {
				fmt.Println(global.GetStyledWarn(errContent))
			}
			break
		}
		if !global.UInput.IsChildProcess {
			break
		}
		global.WaitNextFrame(conf.Env.TickPerSec)
	}
}

func GetCommands() []string {
	return global.Commands
}

func AddCommands(command string) bool {
	global.Commands = append(global.Commands, command)
	return true
}

func PopCommands(num int) bool {
	global.Commands = global.Commands[0 : len(global.Commands)-num]
	return true
}

func ClearCommands() bool {
	global.Commands = []string{}
	return true
}

func InteractiveExecuteCli() string {
	defer fmt.Println(global.GetStyledSystemComent(Exit))
	fmt.Println(global.GetStyledSystemComent(PressEnterToBackToChat))
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
	go handleScpInteractiveInput(&stdinPipe)
	err = shell.Run()
	global.UInput.IsChildProcess = false
	if err != nil {
		return err.Error()
	}
	return buffer.String()
}

func NoInteractiveExecuteCli() string {
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
	go handleScpExit(&stopFlag, shell)
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

func InteractiveExecuteWeb() string {
	var (
		err       error
		shell     *exec.Cmd
		buffer    bytes.Buffer
		stdinPipe io.WriteCloser
		reader    *io.PipeReader
		writer    *io.PipeWriter
	)
	shell = buildShell()
	if stdinPipe, err = shell.StdinPipe(); err != nil {
		return err.Error()
	}
	global.UInput.IsChildProcess = true
	reader, writer = io.Pipe()
	defer reader.Close()
	defer writer.Close()
	shell.Stdout = io.MultiWriter(os.Stdout, writer, &buffer)
	shell.Stderr = io.MultiWriter(os.Stderr, writer, &buffer)
	go stopHook(shell)
	go handleScpInteractiveInput(&stdinPipe)
	go handleScpOutputForWeb(reader)
	err = shell.Run()
	global.UInput.IsChildProcess = false
	if err != nil {
		return err.Error()
	}
	return buffer.String()
}

func NoInteractiveExecuteWeb() string {
	var (
		stopFlag bool
		err      error
		shell    *exec.Cmd
		buffer   bytes.Buffer
		reader   *io.PipeReader
		writer   *io.PipeWriter
	)

	stopFlag = false
	shell = buildShell()
	global.UInput.IsChildProcess = true
	reader, writer = io.Pipe()
	defer reader.Close()
	defer writer.Close()
	shell.Stdout = io.MultiWriter(os.Stdout, writer, &buffer)
	shell.Stderr = io.MultiWriter(os.Stderr, writer, &buffer)
	go stopHook(shell)
	go handleScpExit(&stopFlag, shell)
	go handleScpOutputForWeb(reader)
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
