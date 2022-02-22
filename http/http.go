// https://github.com/fasthttp/router/tree/master/_examples/basic
package http

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func LaunchServer(rDB string) {
	router := InitRooter(rDB)
	fasthttp.ListenAndServe(":"+fmt.Sprintf("%d", 3334), router.Handler)
}
