// Copyright © 2018-2020 Satinderjit Singh.
//
// See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
// the top-level directory of this distribution for the individual copyright
// holder information and the developer policies on copyright and licensing.
//
// Unless otherwise agreed in a custom licensing agreement, no part of the
// kmdgo software, including this file may be copied, modified, propagated.
// or distributed except according to the terms contained in the LICENSE file
//
// Removal or modification of this copyright notice is prohibited.

package utils

import (
	//"fmt"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	// "fmt"
	"os"
)

// KMDGoVersion records the version of KMD Go package
const KMDGoVersion = `0.1.0`

// AppType holds the smart chain's name
type AppType string

// APIParams holds the set of parameters passed along methods to smart chain RPC API queries
type APIParams []interface{}

// APIResult data type for holding returned results of API queries
type APIResult struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

// APIQuery holds the methods and params which is sent with body of RPC API queries
type APIQuery struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

// Error data type to store/format the errors resturned from smart chain API queries
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewAppType returns the smart chain as variable type AppType
func NewAppType(app AppType) *AppType {
	return &app
}

// APICall executes the blockchain RPC Calls and returns the data in JSON output.
// By default this API call tries to get the RPC information from the .conf file from the set blockchain's default directory on the system based on OS which is being executed on.
// If the local system doesn't has the blockchain running, you can also set this RPC settings in environment variables. The environment variable has to use the the combination of the appName + _RPCURL.
// Shell example:
// export komodo_RPCURL="http://192.168.1.1:"
// export komodo_RPCUSER="username"
// export komodo_RPCPASS="password"
// export komodo_RPCPORT="7771"
//
// These environment variables can also be set within Go code. Example:
// os.Setenv("komodo" + "_RPCURL", `http://127.0.0.1:`)
// os.Setenv("komodo" + "_RPCUSER", `username`)
// os.Setenv("komodo" + "_RPCPASS", `password`)
// os.Setenv("komodo" + "_RPCPORT", `7771`)
//
// As per this example, if for example using different SmartChain like "DEX", and the appName is set to example "DEX", just replace it word `komodo` with `DEX`.
func (appMeta AppMetaData) APICall(q *APIQuery) string {
	// fmt.Println("APICall appMeta", appMeta)
	if appMeta.Network == "KMD" || appMeta.Network == "Komodo" {
		appMeta.Network = "komodo"
		// fmt.Println(appMeta.Network)
	}
	var rpcurl, rpcuser, rpcpass, rpcport string

	key := string(appMeta.Network) + "_RPCURL"
	// fmt.Println(key)
	_, ok := os.LookupEnv(key)
	if !ok {
		// fmt.Printf("%s not set\n", key)
		// Try to get RPC info from local default directory for set blockchain
		rpcuser, rpcpass, rpcport = AppRPCInfo(appMeta)
		rpcurl = `http://127.0.0.1:`
	} else {
		// fmt.Printf("%s=%s\n", key, val)
		rpcurl = os.Getenv(string(appMeta.Network) + "_RPCURL")
		rpcuser = os.Getenv(string(appMeta.Network) + "_RPCUSER")
		rpcpass = os.Getenv(string(appMeta.Network) + "_RPCPASS")
		rpcport = os.Getenv(string(appMeta.Network) + "_RPCPORT")
	}

	// fmt.Printf(" %s\n %s\n %s\n %s\n", rpcurl, rpcuser, rpcpass, rpcport)
	// fmt.Printf(" %T\n %T\n %T\n %T\n", rpcurl, rpcuser, rpcpass, rpcport)

	if rpcuser == "" && rpcpass == "" && rpcport == "" {
		// fmt.Println("EMPTY RPC INFO!")
		fmt.Printf("Unable to read or find RPC info for the Network/Blockchain being used by Khoji.\nPlease make sure you have configured config.ini properly.\nIf you still having issues, please contact Khoji developers.\n")
		os.Exit(1)
		// return "EMPTY RPC INFO!"
	}

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	url := rpcurl + rpcport
	// fmt.Println(url)

	var queryStr string
	queryStr = `{"jsonrpc": "1.0", "id":"kmdgo", "method": "` + q.Method + `", "params": ` + q.Params + ` }`
	// fmt.Println("queryStr: ", queryStr)

	queryByte := []byte(queryStr)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(queryByte))
	req.Header.Set("Content-Type", "application/json")

	//req, err := http.NewRequest("POST", , nil)
	req.SetBasicAuth(rpcuser, rpcpass)
	resp, err := client.Do(req)

	// if err != nil {
	// 	log.Println(resp)
	// 	log.Fatalf("==> Error reading API response body for - %v: %v", appName, err)
	// }

	if err != nil {
		// log.Println("resp:", resp)
		// log.Fatalf("==> Error reading API response body for - %v: %v", appName, err)

		// Check if the error is "connection refused", means if the daemon is running or not running or inaccessible due to any X reason.
		matched, _ := regexp.Match(`connection refused`, []byte(err.Error()))
		// fmt.Println(matched, err)
		if matched == true {
			s := string(`{"result": {}, "error": {"code":0,"message":"connection refused"}, "id":"kmdgo"}`)
			// fmt.Println(s)
			return s
		}

		matched, _ = regexp.Match(`context deadline exceeded`, []byte(err.Error()))
		if matched == true {
			s := string(`{"result": {}, "error": {"code":0,"message":"context deadline exceeded"}, "id":"kmdgo"}`)
			// fmt.Println(s)
			return s
		}

		return string(`{"result": {}, "error": {"code":0,"message":"` + err.Error() + `"}, "id":"kmdgo"}`)
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(resp)
		log.Fatalf("==> Error reading API response body for - %v: %v", appMeta.Network, err)
	}

	if len(bodyText) == 0 {
		fmt.Println("SEEMS RETURNED DATA IS EMPTY")
		return `{
			"result": {},
			"error": "EMPTY DATA",
			"id": "curltest"
		  }`
	}

	// fmt.Println("bodyText: ", string(bodyText))
	// fmt.Println("resp.Body: ", resp.Body)
	// fmt.Println("bodyText Bytes: ", bodyText)
	// rspbody := bodyText
	// err2 := json.NewDecoder(resp.Body).Decode(&rspbody)
	// fmt.Println("err2: ", err2)
	// fmt.Println("rspbody: ", rspbody)

	var queryResult map[string]interface{}
	if err := json.Unmarshal(bodyText, &queryResult); err != nil {
		fmt.Println("bodyText on json marshel: ", string(bodyText))
		fmt.Println("resp.Body: ", resp.Body)
		fmt.Println("bodyText Bytes: ", bodyText)
		panic(err)
	}

	s := string(bodyText)
	return s
}

// RPCResultMap using golang's own http package
func (appMeta AppMetaData) RPCResultMap(method string, params interface{}) (interface{}, error) {

	// fmt.Println("RPCResultMap appMeta", appMeta)

	// fmt.Printf("params -- %+v\n", params)

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println("paramsJSON:", string(paramsJSON))

	query := APIQuery{
		Method: method,
		Params: string(paramsJSON),
	}
	// fmt.Printf("%+v\n", query)

	getJSON := appMeta.APICall(&query)
	if getJSON == "EMPTY RPC INFO" {
		fmt.Println("EMPTY RPC INFO")
		return nil, errors.New("EMPTY RPC INFO")
	}
	// fmt.Printf("getJSON - %v\n", getJSON)

	var result map[string]interface{}
	json.Unmarshal([]byte(getJSON), &result)
	// fmt.Println(result["error"])
	if result["error"] != nil {
		if result["error"].(map[string]interface{})["message"].(string) == "connection refused" {
			fmt.Println("ERROR: connection refused. Khoji is unable to connect to the blockchain API.\nPlease check:\n- The blockchain daemon is runnning and it is configured to accept API connections from external programs.\n- Make sure config.ini file of Khoji explorer is configured properly if using remote blockchain node instead of local blockchain API.")
			os.Exit(1)
		}
		if result["error"].(map[string]interface{})["message"].(string) == "Loading block index..." {
			fmt.Println("ERROR: Loading block index... Looks like the blockchain daemon is still loading.\nPlease wait for it to load completely and be ready to accept connections from external programs\nand try executing Khoji to connect with it again in a momnent.")
			os.Exit(1)
		}
		// fmt.Printf("%v\n", result["error"].(map[string]interface{})["message"])
		return nil, errors.New(result["error"].(map[string]interface{})["message"].(string))
	}
	return result["result"], nil
	// return "", nil
}
