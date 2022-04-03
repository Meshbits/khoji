// https://github.com/fasthttp/router/tree/master/_examples/basic
package http

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
	"gopkg.in/ini.v1"
)

// ref: https://github.com/valyala/fasthttp/blob/master/fs_handler_example_test.go
// Setup file handlers (aka 'file server config')
var (
	// Handler for serving images from /img/ path,
	// i.e. /img/foo/bar.jpg will be served from
	// /var/www/images/foo/bar.jpb .
	imgPrefix  = []byte("/img/")
	imgHandler = fasthttp.FSHandler("/var/www/images", 1)

	// Handler for serving css from /static/css/ path,
	// i.e. /static/css/foo/bar.css will be served from
	// /home/dev/css/foo/bar.css .
	cssPrefix  = []byte("/static/css/")
	cssHandler = fasthttp.FSHandler("/home/dev/css", 2)

	// Handler for serving the rest of requests,
	// i.e. /foo/bar/baz.html will be served from
	// /var/www/files/foo/bar/baz.html .
	filesHandler = fasthttp.FSHandler("/var/www/files", 0)
)

// Main request handler
func requestHandler(ctx *fasthttp.RequestCtx) {
	path := ctx.Path()
	switch {
	case bytes.HasPrefix(path, imgPrefix):
		imgHandler(ctx)
	case bytes.HasPrefix(path, cssPrefix):
		cssHandler(ctx)
	default:
		filesHandler(ctx)
	}
}

func LaunchServer() {
	router := InitRooter()

	cfg, err := ini.ShadowLoad("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	autoSSL, _ := cfg.Section("SSL").Key("ENABLE_AUTOCERT").Bool()
	fmt.Println("autoSSL", autoSSL)

	hosts := cfg.Section("SSL").Key("HOST").ValueWithShadows()
	// fmt.Printf("hosts: %q\n", hosts)

	if autoSSL {
		m := &autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(hosts...), // Replace with your domain.
			Cache:      autocert.DirCache("./certs"),
		}

		tlsCfg := &tls.Config{
			GetCertificate: m.GetCertificate,
			NextProtos: []string{
				"http/1.1", acme.ALPNProto,
			},
		}

		// Let's Encrypt tls-alpn-01 only works on port 443.
		ln, err := net.Listen("tcp4", "0.0.0.0:443") /* #nosec G102 */
		if err != nil {
			panic(err)
		}

		lnTls := tls.NewListener(ln, tlsCfg)
		for _, v := range hosts {
			fmt.Println("-------------------------------")
			fmt.Printf("Web UI: https://%v\n", v)
			fmt.Printf("API: https://%v/api\n", v)
			fmt.Println("-------------------------------")
		}
		fmt.Println("")

		if err := fasthttp.Serve(lnTls, router.Handler); err != nil {
			panic(err)
		}
	} else {
		fmt.Println("-------------------------------")
		fmt.Printf("Web UI: http://localhost:3334\n")
		fmt.Printf("API: http://localhost:3334/api\n")
		fmt.Println("-------------------------------")
		fasthttp.ListenAndServe(":"+fmt.Sprintf("%d", 3334), router.Handler)
	}

}
