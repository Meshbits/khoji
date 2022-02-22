package http

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func LaunchServer() {
	router := InitRooter()
	fasthttp.ListenAndServe(":"+fmt.Sprintf("%d", 3334), router.Handler)
}
