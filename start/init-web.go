package start

import (
	"asashishi-agent/global"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func WithWebMode() {
	var (
		err        error
		dirPath    string
		fileServer http.Handler
	)
	if dirPath, err = os.Getwd(); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
	fileServer = http.FileServer(http.Dir(filepath.Join(dirPath, global.WebServerRootPath)))
	http.Handle("/", fileServer)

	fmt.Println(
		global.GetStyledSuccess(
			fmt.Sprintf(global.WebServerStartComment, global.WebServerPort),
		),
	)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", global.WebServerPort), nil); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
}
