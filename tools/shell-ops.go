package tools

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"bufio"
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
		fomatedCommands string
		shell           *exec.Cmd
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
	go func() {
		for {
			select {
			case <-global.UInput.ChildProcessStdin:
				stopFlag = true
				killChildProcessGroup(shell.Process.Pid)
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

// web 模式待修改 使用 channel 将 io.MultiWriter 的输出接到另一个无缓冲管道

func InteractiveExecuteWeb() string {
	var (
		err       error
		shell     *exec.Cmd
		buffer    *bytes.Buffer
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
	defer writer.Close()
	defer reader.Close()
	shell.Stdout = io.MultiWriter(os.Stdout, writer)
	shell.Stderr = io.MultiWriter(os.Stderr, writer)
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
	go func() {
		var (
			innerErr error
			line     string
			scanner  *bufio.Scanner
		)
		scanner = bufio.NewScanner(reader)
		for scanner.Scan() {
			line = scanner.Text()
			global.ScpOutputChan <- line
		}
		if innerErr = scanner.Err(); innerErr != nil {
			fmt.Println(global.GetStyledWarn(innerErr.Error()))
		}
	}()
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
		buffer   *bytes.Buffer
		reader   *io.PipeReader
		writer   *io.PipeWriter
	)

	stopFlag = false
	shell = buildShell()
	global.UInput.IsChildProcess = true

	reader, writer = io.Pipe()
	defer writer.Close()
	defer reader.Close()
	shell.Stdout = io.MultiWriter(os.Stdout, writer)
	shell.Stderr = io.MultiWriter(os.Stderr, writer)
	go func() {
		for {
			select {
			case <-global.UInput.ChildProcessStdin:
				stopFlag = true
				killChildProcessGroup(shell.Process.Pid)
			default:
				if !global.UInput.IsChildProcess {
					break
				}
				global.WaitNextFrame(conf.Env.TickPerSec)
			}
		}
	}()
	go func() {
		var (
			innerErr error
			line     string
			scanner  *bufio.Scanner
		)
		scanner = bufio.NewScanner(reader)
		for scanner.Scan() {
			line = scanner.Text()
			global.ScpOutputChan <- line
		}
		if innerErr = scanner.Err(); innerErr != nil {
			fmt.Println(global.GetStyledWarn(innerErr.Error()))
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
