package main

import (
	"fmt"

	"github.com/Meshbits/khoji/utils"
)

// Define appName type from kmdgo package
// Define appname variable. The name value must be the matching value of it's data directory name.
// Example Komodo's data directory is `komodo`, VerusCoin's data directory is `VRSC` and so on.
var appName utils.AppType

// Blockchain
var ChainName string

func init() {
	ChainName = `VRSCTEST`
}

func main() {
	appDir := utils.AppDataDir(ChainName, false)
	fmt.Println("appDir:", appDir)

	appName = utils.AppType(ChainName)
	// fmt.Println("appName:", appName)

	// `{"systemtype":"pbaas"}`
	// Collect listCurrencies information
	// _listCurrencies, _ := appName.RPCResultMap("listcurrencies", []interface{}{map[string]string{"systemtype": "pbaas"}})
	// // fmt.Println("_listCurrencies", _listCurrencies)
	// listCurrencies := _listCurrencies.([]interface{})
	// fmt.Println("listCurrencies length:", len(listCurrencies))
	// fmt.Println("listCurrencies:", listCurrencies[0])
}
