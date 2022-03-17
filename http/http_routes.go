package http

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/Meshbits/khoji/shepherd"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"gopkg.in/ini.v1"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type respBalanceOk struct {
	Balance float64 `result:"balance"`
}
type respErr struct {
	Error string `result:"error"`
}

var MAX_ITEMS_PP int = 10

var session *r.Session

// Rethink database name
var rDB string

func init() {
	// fmt.Println("http_routes")

	var err error
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// rDB = os.Getenv("RDB_DB")
	rDB = cfg.Section("DATABASE").Key("RDB_DB").String()
	session, err = r.Connect(r.ConnectOpts{
		Address: cfg.Section("DATABASE").Key("RDB_IP").String() + ":" + cfg.Section("DATABASE").Key("RDB_PORT").String(),
		// Database: rDB,
	})
	if err != nil {
		fmt.Printf("ERROR: There is issue connecting with the database.\nPlease make sure databse is accessible to Khoji by making sure settings in\nconfig.ini are setup properly and the database server is up and running.\n\n")
		fmt.Println("ERROR DETAILS:", err)
		os.Exit(1)
		return
	}
}

func setResponseHeader(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		h(ctx)
	}
}

func getAddressBalance(ctx *fasthttp.RequestCtx) {
	address := ctx.UserValue("address").(string)
	res1, err1 := r.DB(rDB).Table("accounts").Filter(map[string]interface{}{"address": address}).Map(
		func(row r.Term) interface{} { return row.Field("balance") }).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get balance info from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row interface{}
	err2 := res1.One(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	if row != nil {
		fmt.Println("row", row)
		jsonData, _ := json.Marshal(respBalanceOk{
			Balance: row.(float64),
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr{
			Error: "No such address",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
		fmt.Println("address", address)
	}
}

func getAddressTransactions(ctx *fasthttp.RequestCtx) {
	address := ctx.UserValue("address").(string)

	res1, err1 := r.DB(rDB).Table("accounts").Filter(map[string]interface{}{"address": address}).Map(
		func(row r.Term) interface{} { return row.Field("transactions") }).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get address transactions from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row interface{}
	err2 := res1.One(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	if row != nil {
		var txDetails []interface{}

		for _, txid := range row.([]interface{}) {
			log.Printf("txid %v", txid)

			res3, err3 := r.DB(rDB).Table("transactions").Filter(map[string]interface{}{"hash": txid}).Run(session)
			if err3 != nil {
				log.Panicf("Failed to get transaction info from DB: %v", err3)
			}
			log.Printf("query res %v", res3)
			var row2 []interface{}
			err4 := res3.All(&row2)
			if err4 == r.ErrEmptyResult {
				fmt.Println("row not found")
			}
			if err4 != nil {
				fmt.Println(err2)
			}

			fmt.Println("row", row2)
			//fmt.Println("row", row[0]["height"])

			if row2 != nil {
				txDetails = append(txDetails, row2[0])
			}
		}

		fmt.Println("row", row)
		jsonData, _ := json.Marshal(txDetails)
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr{
			Error: "No such address",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
		fmt.Println("address", address)
	}
}

func getBlockInfo(ctx *fasthttp.RequestCtx) {
	height := ctx.UserValue("height").(string)
	heightInt, _ := strconv.Atoi(height)

	res1, err1 := r.DB(rDB).Table("blocks").Filter(map[string]interface{}{"height": heightInt}).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get block info from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
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
		jsonData, _ := json.Marshal(respErr{
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

	log.Printf("get blocks from: %v to %v", pageInt*MAX_ITEMS_PP, (pageInt+1)*MAX_ITEMS_PP)

	res1, err1 := r.DB(rDB).Table("blocks").Without("transactions", "solution").OrderBy("height").Slice(pageInt*MAX_ITEMS_PP, (pageInt+1)*MAX_ITEMS_PP).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get block info from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
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
		jsonData, _ := json.Marshal(respErr{
			Error: "Wrong page number",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func getTransactionDetails(ctx *fasthttp.RequestCtx) {
	hash := ctx.UserValue("hash").(string)

	res1, err1 := r.DB(rDB).Table("transactions").Filter(map[string]interface{}{"hash": hash}).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get transaction details from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("txid", hash)

	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(row[0])
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr{
			Error: "No such transaction",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func getIdentityDetails(ctx *fasthttp.RequestCtx) {
	name := ctx.UserValue("name").(string)

	res1, err1 := r.DB(rDB).Table("identities").Filter(map[string]interface{}{"name": name}).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get identity details from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("name", name)

	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(row[0])
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr{
			Error: "No such identity",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func getIdentitiesSlice(ctx *fasthttp.RequestCtx) {
	page := ctx.UserValue("page").(string)
	pageInt, _ := strconv.Atoi(page)

	log.Printf("get identities from: %v to %v", pageInt*MAX_ITEMS_PP, (pageInt+1)*MAX_ITEMS_PP)

	res1, err1 := r.DB(rDB).Table("identities").OrderBy("blockheight").Slice(pageInt*MAX_ITEMS_PP, (pageInt+1)*MAX_ITEMS_PP).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get identities info from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
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
		jsonData, _ := json.Marshal(respErr{
			Error: "Wrong page number",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func getNetworkInfo(ctx *fasthttp.RequestCtx) {
	res1, err1 := r.DB(rDB).Table("network").Run(session)
	if err1 != nil {
		log.Panicf("Failed to get network info from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(row[0])
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr{
			Error: "No network data",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

// TODO: address transactions

func getAccountsSlice(ctx *fasthttp.RequestCtx) {
	page := ctx.UserValue("page").(string)
	pageInt, _ := strconv.Atoi(page)

	log.Printf("get accounts from: %v to %v", pageInt*MAX_ITEMS_PP, (pageInt+1)*MAX_ITEMS_PP)

	res1, err1 := r.DB(rDB).Table("accounts").OrderBy(r.Desc("balance")).Without("transactions").Slice(pageInt*MAX_ITEMS_PP, (pageInt+1)*MAX_ITEMS_PP).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get richlist from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
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
		jsonData, _ := json.Marshal(respErr{
			Error: "Wrong page number",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func getLastBlocks(ctx *fasthttp.RequestCtx) {
	var pageInt = 0

	res1, err1 := r.DB(rDB).Table("blocks").Without("transactions", "solution").OrderBy(r.Desc("height")).Slice(pageInt*MAX_ITEMS_PP, (pageInt+1)*MAX_ITEMS_PP).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get last blocks from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(row)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr{
			Error: "Wrong page number",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func getLastTransactions(ctx *fasthttp.RequestCtx) {
	var pageInt = 0

	res1, err1 := r.DB(rDB).Table("transactions").OrderBy(r.Desc("height")).Slice(pageInt*MAX_ITEMS_PP, (pageInt+1)*MAX_ITEMS_PP).Run(session)
	if err1 != nil {
		log.Panicf("Failed to get last transactions from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(row)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr{
			Error: "Wrong page number",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func checkUpdate(ctx *fasthttp.RequestCtx) {
	os := ctx.UserValue("os")
	arch := ctx.UserValue("arch")
	// fmt.Printf("OS: %v\nARCH: %v\n", os, arch)

	if os == nil {
		os = runtime.GOOS
	}
	if arch == nil {
		arch = runtime.GOARCH
	}

	res1, err1 := r.DB(rDB).Table("network").Run(session)
	if err1 != nil {
		log.Panicf("Failed to get network info from DB: %v", err1)
	}
	// log.Printf("query res %v", res1)
	var row []interface{}
	err2 := res1.All(&row)
	if err2 == r.ErrEmptyResult {
		fmt.Println("row not found")
	}
	if err2 != nil {
		fmt.Println(err2)
	}

	currentVersion := "v" + row[0].(map[string]interface{})["VRSCversion"].(string)

	getUpdate := shepherd.GetDlURL(os.(string), arch.(string), currentVersion)
	if getUpdate.Err != nil {
		fmt.Println(getUpdate.Err)
	}
	getUpdate.CurrentVersion = currentVersion

	if row != nil {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(getUpdate)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	} else {
		ctx.SetStatusCode(200)
		jsonData, _ := json.Marshal(respErr{
			Error: "something went wrong",
		})
		ctx.SetStatusCode(200)
		ctx.SetBodyString(string(jsonData))
		ctx.SetContentType("application/json")
	}
}

func InitRooter() *router.Router {
	r := router.New()

	r.GET("/api/v1/network", setResponseHeader(getNetworkInfo))
	r.GET("/api/v1/balance/{address}", setResponseHeader(getAddressBalance))
	r.GET("/api/v1/transactions/{address}", setResponseHeader(getAddressTransactions))
	r.GET("/api/v1/transactions/last", setResponseHeader(getLastTransactions))
	r.GET("/api/v1/transaction/{hash}", setResponseHeader(getTransactionDetails))
	r.GET("/api/v1/block/{height}", setResponseHeader(getBlockInfo))
	r.GET("/api/v1/blocks/{page}", setResponseHeader(getBlocksSlice))
	r.GET("/api/v1/blocks/last", setResponseHeader(getLastBlocks))
	r.GET("/api/v1/identity/{name}", setResponseHeader(getIdentityDetails))
	r.GET("/api/v1/identities/{page}", setResponseHeader(getIdentitiesSlice))
	r.GET("/api/v1/richlist/{page}", setResponseHeader(getAccountsSlice))
	r.GET("/api/v1/checkupdate/{os?}/{arch?}", setResponseHeader(checkUpdate))

	return r
}
