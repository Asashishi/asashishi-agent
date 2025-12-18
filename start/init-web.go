package start

import (
	"asashishi-agent/conf"
	"asashishi-agent/global"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func WithWebMode() {

	InitWeb()

	var (
		err        error
		dirPath    string
		fileServer http.Handler
	)

	if dirPath, err = os.Getwd(); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
	fileServer = http.FileServer(http.Dir(filepath.Join(dirPath, conf.Env.ServerRootPath)))
	http.Handle("/", fileServer)

	fmt.Println(
		global.GetStyledSuccess(
			fmt.Sprintf(global.WebServerStartComment, conf.Env.ServerPort),
		),
	)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", conf.Env.ServerPort), nil); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
}
