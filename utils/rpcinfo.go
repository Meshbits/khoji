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
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

// BytesToString converts and returns bytes to string
func BytesToString(data []byte) string {
	return string(data[:])
}

// AppRPCInfo returns RPC username, password, port info for a specified Komodo Assetchain (Antara smartchain)
func AppRPCInfo(appMeta AppMetaData) (string, string, string) {
	appDir := AppDataDir(appMeta)
	// fmt.Println("appDir", appDir)

	var appConf string

	if strings.ToLower(appMeta.Network) == "komodo" || strings.ToLower(appMeta.Network) == ".komodo" {
		if runtime.GOOS == "darwin" {
			appConf = filepath.Join(appDir, strings.Title(appMeta.Network)+`.conf`)
		} else {
			appConf = filepath.Join(appDir, strings.ToLower(appMeta.Network)+`.conf`)
		}
	} else {
		if strings.ToLower(appMeta.Network) == "vrsctest" {
			if appMeta.IsPBaaS != true {
				appConf = filepath.Join(appDir, strings.ToLower(appMeta.Network)+`.conf`)
			} else {
				appConf = filepath.Join(appDir, strings.ToLower(appMeta.PBaaS)+`.conf`)
			}
		} else {
			appConf = filepath.Join(appDir, strings.ToUpper(appMeta.Network)+`.conf`)
		}
	}

	// fmt.Println("appConf", appConf)

	confdata, err := ioutil.ReadFile(appConf)
	if err != nil {
		switch err {
		case os.ErrInvalid:
			//Do stuff
		case os.ErrPermission:
			//Do stuff
		case os.ErrNotExist:
			// fmt.Println("File Does not exists")
			// log.Println(err)
		default:
			// fmt.Println(err)
		}
		return "", "", ""
	}

	var rpcu = regexp.MustCompile("(?m)^rpcuser=.+$")
	var rpcpass = regexp.MustCompile("(?m)^rpcpassword=.+$")
	var rpcport = regexp.MustCompile("(?m)^rpcport=.+$")

	// fmt.Println("confdata: ", confdata)

	bytestr := BytesToString(confdata)

	rpcuserLine := rpcu.FindString(bytestr)
	rpcpassLine := rpcpass.FindString(bytestr)
	rpcpportLine := rpcport.FindString(bytestr)

	// fmt.Println(bytestr)

	//AppRPCuser := strings.TrimLeft(strings.TrimLeft(rpcuserLine,`rpcuser`),`=`)
	//AppRPCpass := strings.TrimLeft(strings.TrimLeft(rpcpassLine,`rpcpassword`),`=`)
	AppRPCport := strings.TrimLeft(strings.TrimLeft(rpcpportLine, `rpcport`), `=`)

	//AppRPCuser := rpcuserLine[8:]
	AppRPCuser := strings.TrimLeft(rpcuserLine, `rpcuser`)[1:]
	AppRPCpass := strings.TrimLeft(rpcpassLine, `rpcpassword`)[1:]
	//AppRPCport := strings.TrimLeft(rpcpportLine,`rpcport`)[1:]

	if AppRPCport == "" && appMeta.Network == "komodo" {
		AppRPCport = "7771"
	}

	//TODO
	// Check if server=1 and update/change it
	// Check if rpcuser and rpcpassword is the known default value from QT wallet,
	// example "komodo" and "local321", and prompt/update it with random values.

	// Removing any \n or \r\n from the stripped string from .conf file
	AppRPCuser = strings.ReplaceAll(AppRPCuser, "\n", "")
	AppRPCuser = strings.ReplaceAll(AppRPCuser, "\r", "")
	AppRPCpass = strings.ReplaceAll(AppRPCpass, "\n", "")
	AppRPCpass = strings.ReplaceAll(AppRPCpass, "\r", "")
	AppRPCport = strings.ReplaceAll(AppRPCport, "\n", "")
	AppRPCport = strings.ReplaceAll(AppRPCport, "\r", "")
	return AppRPCuser, AppRPCpass, AppRPCport
}
