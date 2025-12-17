package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

var Env EnvConfig = EnvConfig{}

func InitConfig() {
	var (
		err     error
		dirName string
		file    *os.File
		decoder *json.Decoder
	)
	if EnvDetect() == Windows {
		Env.System = Windows
		Env.SysPrompt = fmt.Sprintf(SysPrompt, Windows, PowerShell)
	} else {
		Env.System = Linux
		Env.SysPrompt = fmt.Sprintf(SysPrompt, Linux, Bash)
	}

	if dirName, err = os.Getwd(); err != nil {
		panic(err)
	} else if file, err = os.Open(path.Join(dirName, ConfigJsonName)); err != nil {
		panic(err)
	}
	defer file.Close()
	decoder = json.NewDecoder(file)
	if err = decoder.Decode(&Env); err != nil {
		panic(err)
	}
	Env.Version = Version
}
