package tools

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"sync"
)

var Commands []string = []string{}

func killChildProcessGroup(pid string) {
	exec.Command("taskkill", "/PID", pid, "/F", "/T").Run()
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

func Execute() string {
	fmt.Println("\n-- Press 'enter' to back to chat")
	defer fmt.Println("\n-- Exit")
	var (
		stopFlag        bool
		err             error
		fomatedCommands string
		shell           *exec.Cmd
		buffer          bytes.Buffer
		wg              sync.WaitGroup
	)
	stopFlag = false
	fomatedCommands = InitialCommand
	for _, cmd := range Commands {
		fomatedCommands = fmt.Sprintf("%s; if ($?) { %s }", fomatedCommands, cmd)
	}
	Commands = []string{}
	shell = exec.Command("powershell", "-Command", fomatedCommands)
	shell.Stdin = os.Stdin
	shell.Stdout = io.MultiWriter(os.Stdout, &buffer)
	shell.Stderr = io.MultiWriter(os.Stderr, &buffer)
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		var (
			innerErr error
			reader   *bufio.Reader
		)
		reader = bufio.NewReader(os.Stdin)
		if _, innerErr = reader.ReadString('\n'); innerErr != nil {
			stopFlag = true
		}
		stopFlag = true
		killChildProcessGroup(strconv.Itoa(shell.Process.Pid))
	}()
	if err = shell.Run(); err != nil {
		if stopFlag {
			return buffer.String() + StopedByUser
		}
		return err.Error()
	}
	return buffer.String()
}
