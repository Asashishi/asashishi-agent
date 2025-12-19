package global

import (
	"asashishi-agent/conf"
	"asashishi-agent/ui"
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var Commands []string = []string{}
var UInput *GlobalUInput = &GlobalUInput{
	IsChildProcess:    false,
	ProcessStdin:      make(chan string),
	ChildProcessStdin: make(chan string),
}
var ScpOutputChan chan string = make(chan string)

func InitGlobalCliUInput() {
	var (
		err    error
		str    string
		reader *bufio.Reader
	)
	reader = bufio.NewReader(os.Stdin)
	for {
		if str, err = reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				fmt.Println()
				os.Exit(0)
			}
			panic(GetStyledError(err.Error()))
		}
		if UInput.IsChildProcess {
			UInput.ChildProcessStdin <- str
		} else {
			UInput.ProcessStdin <- str
		}
	}
}

func InitGlobalWebUInput() {
	var str string
	UInput.WebsocketReadChan = make(chan string)
	for {
		select {
		case str = <-UInput.WebsocketReadChan:
			if UInput.IsChildProcess {
				UInput.ChildProcessStdin <- str
			} else {
				UInput.ProcessStdin <- str
			}
		default:
			WaitNextFrame(conf.Env.TickPerSec)
		}
	}
}

func WaitNextFrame(tick float64) {
	time.Sleep(
		(time.Duration((FloatK / tick) * FloatK)) * time.Microsecond,
	)
}

func SetTerminalTitle() {
	fmt.Printf(TilteEscape, AppTitle)
}

func PrintAppBanner(version string) {
	var styledVersion string = fmt.Sprintf(Version, version)
	fmt.Printf(AppBanner, styledVersion)
}

func GetStyledWarn(warnString string) string {
	return ui.WidthStyle(
		warnString,
		SystemWarnStyle,
	)
}

func GetStyledError(errString string) string {
	return ui.WidthStyle(
		errString,
		SystemErrorStyle,
	)
}

func GetStyledSuccess(sucStirng string) string {
	return ui.WidthStyle(
		sucStirng,
		SystemSuccesStyle,
	)
}

func GetStyledSystemComent(comentString string) string {
	return ui.WidthStyle(
		comentString,
		SystemCommentStyle,
	)
}
