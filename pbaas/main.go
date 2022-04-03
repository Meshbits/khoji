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
	// Just switch IsPBaaS value from 'true' to 'false' to get different results
	// IsPBaaS: true, will give RPC info of PBaaS chain with IDHex defined in PBaaS
	// IsPBaaS: false, will give just the ChainName RPC info despite PBaaS IDHex defined
	appMeta := utils.AppMetaData{
		Network: ChainName,
		IsPBaaS: true,
		PBaaS:   `3e76382e8354715b3f0be56608c112174baaf554`,
	}
	if appMeta.IsPBaaS == true {
		fmt.Println("IsPBaaS:", appMeta.IsPBaaS)
		fmt.Println("Returning data directory and RPC info for PBaaS chain.")
		fmt.Println("======================================================")
	} else {
		fmt.Println("IsPBaaS:", appMeta.IsPBaaS)
		fmt.Printf("Returning data directory and RPC info for %v network.\n", appMeta.Network)
		fmt.Println("======================================================")
	}
	appDir := utils.AppDataDir(appMeta)
	fmt.Println("appDir:", appDir)

	rpcuser, rpcpass, rpcport := utils.AppRPCInfo(appMeta)
	fmt.Printf("rpcuser: %s\nrpcpass: %s\nrpcport: %s\n", rpcuser, rpcpass, rpcport)

	// `{"systemtype":"pbaas"}`
	// Collect listCurrencies information
	_listCurrencies, _ := appMeta.RPCResultMap("listcurrencies", []interface{}{map[string]string{"systemtype": "pbaas"}})
	// fmt.Println("_listCurrencies", _listCurrencies)
	listCurrencies := _listCurrencies.([]interface{})

	fmt.Println("--------")
	// List all PBaaS currencies/chains names
	fmt.Println("listCurrencies length:", len(listCurrencies))
	// fmt.Println("listCurrencies:", listCurrencies[0].(map[string]interface{})["currencydefinition"].(map[string]interface{})["name"])
	for _, v := range listCurrencies {
		if v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["name"] != "vrsctest" {
			fmt.Printf("name:id_hex - %v:%v\n", v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["name"], v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["currencyidhex"])
		}
	}

	fmt.Println("--------")
	// Only list specific currency
	fmt.Println("listCurrencies:", listCurrencies[0].(map[string]interface{})["currencydefinition"].(map[string]interface{})["name"])
	for _, v := range listCurrencies {
		if v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["name"] == `PBaaS` {
			fmt.Printf("name:id_hex - %v:%v\n", v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["name"], v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["currencyidhex"])
		}
	}
}
