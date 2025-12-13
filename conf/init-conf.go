package conf

import (
	"encoding/json"
	"os"
	"path"
)

var Env EnvConfig = EnvConfig{}

func InitConfig() {
	var (
		err       error
		dirName   string
		sysPrompt []byte
		file      *os.File
		decoder   *json.Decoder
	)
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
	// prompt
	if sysPrompt, err = os.ReadFile(path.Join(dirName, SysPromptName)); err != nil {
		panic(err)
	}
	Env.SysPrompt = string(sysPrompt)
}
