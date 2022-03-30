package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	fs := &fasthttp.FS{
		Root:               "./dist/",
		IndexNames:         []string{"index.html"},
		GenerateIndexPages: false,
		Compress:           false,
		AcceptByteRange:    false,
	}
	fsHandler := fs.NewRequestHandler()

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fsHandler(ctx)
	}

	fasthttp.ListenAndServe(":"+fmt.Sprintf("%d", 9999), requestHandler)
}
