package start

import (
	"asashishi-agent/global"
	"fmt"
	"net/http"
)

func WithWebMode() {
	var (
		err        error
		fileServer http.Handler
	)
	fileServer = http.FileServer(http.Dir("../web"))
	http.Handle("/", fileServer)

	fmt.Println(
		global.GetStyledSuccess(
			fmt.Sprintf(WebServerStartComment, WebServerPort),
		),
	)

	if err = http.ListenAndServe(fmt.Sprintf(":%d", WebServerPort), nil); err != nil {
		panic(global.GetStyledError(err.Error()))
	}
}
