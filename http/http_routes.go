// ref: https://github.com/KomodoPlatform/etherscan-mm2-proxy/blob/master/http/http_routes.go

package http

import (
	//"strings"
	"fmt"
	"log"
	"encoding/json"
	"strconv"
	rdb "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type respBalanceOk struct {
	Balance float64 `result:"balance"`
}
type respErr struct {
	Error string `result:"error"`
}

// session for rethink db
var session *rdb.Session
var rDB string
var MAX_ITEMS_PP int = 10

func init() {
	rDB = "vrsctest"

	var err error
	session, err = rdb.Connect(rdb.ConnectOpts{
		Address: "localhost:28015",
		// Database: rDB,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setResponseHeader(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		h(ctx)
		return
	}
}

func getNetworkInfo(ctx *fasthttp.RequestCtx) {	
	res1, err1 := rdb.DB(rDB).Table("network").Run(session)
	if err1 != nil {
		log.Panicf("Failed to get network info from DB: %v", err1)
	}
	log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == rdb.ErrEmptyResult {
		// row not found
	}
	if err2 != nil {
		// error
	}
	
	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(row[0])
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr {
			Error: "No network data",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func getBlockInfo(ctx *fasthttp.RequestCtx) {
	height := ctx.UserValue("height").(string)
	heightInt, _ := strconv.Atoi(height)
	
	res1, err1 := rdb.DB(rDB).Table("blocks").Filter(map[string]interface{}{"height": heightInt}).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get block info from DB: %v", err1)
	}
	log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == rdb.ErrEmptyResult {
		// row not found
	}
	if err2 != nil {
		// error
	}

	fmt.Println("height", height)	
	//fmt.Println("row", row[0]["height"])
	
	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(row[0])
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr {
			Error: "No such block",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func getBlocksSlice(ctx *fasthttp.RequestCtx) {
	page := ctx.UserValue("page").(string)
	pageInt, _ := strconv.Atoi(page)

	log.Printf("get blocks from: %v to %v", pageInt * MAX_ITEMS_PP, (pageInt + 1) * MAX_ITEMS_PP)

	res1, err1 := rdb.DB(rDB).Table("blocks").Without("transactions", "solution").OrderBy("height").Slice(pageInt * MAX_ITEMS_PP, (pageInt + 1) * MAX_ITEMS_PP).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get block info from DB: %v", err1)
	}
	log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == rdb.ErrEmptyResult {
		// row not found
	}
	if err2 != nil {
		// error
	}

	fmt.Println("page", page)	
	//fmt.Println("row", row[0]["height"])
	
	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(row)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr {
			Error: "Wrong page number",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func getLastBlocks(ctx *fasthttp.RequestCtx) {
	var pageInt = 0

	res1, err1 := rdb.DB(rDB).Table("blocks").Without("transactions", "solution").OrderBy(rdb.Desc("height")).Slice(pageInt * MAX_ITEMS_PP, (pageInt + 1) * MAX_ITEMS_PP).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get last blocks from DB: %v", err1)
	}
	log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == rdb.ErrEmptyResult {
		// row not found
	}
	if err2 != nil {
		// error
	}
	
	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(row)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr {
			Error: "Wrong page number",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}


func InitRooter() *router.Router {
	r := router.New()

	r.GET("/api/network", setResponseHeader(getNetworkInfo))
	r.GET("/api/block/{height}", setResponseHeader(getBlockInfo))
	r.GET("/api/blocks/{page}", setResponseHeader(getBlocksSlice))
	r.GET("/api/blocks/last", setResponseHeader(getLastBlocks))

	return r
}