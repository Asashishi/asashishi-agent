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
	"strings"
	"sync"
)

var Commands []string = []string{}

func readProcessChildren(pid int) []int {
	var (
		err      error
		data     []byte
		path     string
		textData []string
		children []int = []int{}
	)
	path = fmt.Sprintf("/proc/%d/task/%d/children", pid, pid)
	if data, err = os.ReadFile(path); err != nil {
		fmt.Println(GetStyledError(err.Error()))
		return children
	}
	if textData = strings.Split(strings.TrimSpace(string(data)), global.SpaceString); len(textData) == 0 {
		return children
	}
	for _, s := range textData {
		if s == "" {
			continue
		}
		if cpid, err := strconv.Atoi(s); err == nil {
			children = append(children, cpid)
		} else {
			fmt.Println(GetStyledError(err.Error()))
		}
	}
	return children
}

func killProcessTree(pid int) {
	var childrens []int = readProcessChildren(pid)
	for _, c := range childrens {
		killProcessTree(c)
	}
	exec.Command("kill", "-9", strconv.Itoa(pid)).Run()
}

func killChildProcessGroup(pid int) {
	if conf.Env.System == conf.Windows {
		exec.Command("taskkill", "/PID", strconv.Itoa(pid), "/F", "/T").Run()
	} else {
		killProcessTree(pid)
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
		err    error
		shell  *exec.Cmd
		buffer bytes.Buffer
	)
	shell = buildShell()
	shell.Stdin = os.Stdin
	shell.Stdout = io.MultiWriter(os.Stdout, &buffer)
	shell.Stderr = io.MultiWriter(os.Stderr, &buffer)
	if err = shell.Run(); err != nil {
		return err.Error()
	}
	return buffer.String()
}

func NoInterActiveExecute() string {
	fmt.Println(PressEnterToBackToChat)
	defer fmt.Println(Exit)
	var (
		stopFlag bool
		err      error
		shell    *exec.Cmd
		buffer   bytes.Buffer
		wg       sync.WaitGroup
	)
	stopFlag = false
	shell = buildShell()
	shell.Stdout = io.MultiWriter(os.Stdout, &buffer)
	shell.Stderr = io.MultiWriter(os.Stderr, &buffer)
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		var reader *bufio.Reader = bufio.NewReader(os.Stdin)
		reader.ReadString(global.LineBreakChar)
		stopFlag = true
		killChildProcessGroup(shell.Process.Pid)
	}()
	if err = shell.Run(); err != nil {
		if stopFlag {
			return buffer.String() + StopedByUser
		}
		return err.Error()
	}
	return buffer.String()
}
