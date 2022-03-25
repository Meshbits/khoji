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
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

type AppMetaData struct {
	Goos, PBaaS, Network string
	Roaming, IsPBaaS     bool
}

// appDataDir returns an operating system specific directory to be used for
// storing application data for an application.  See AppDataDir for more
// details.  This unexported version takes an operating system argument
// primarily to enable the testing package to properly test the function by
// forcing an operating system that is not the currently one.
func appDataDir(appMeta AppMetaData) string {
	if appMeta.Goos == "" {
		appMeta.Goos = runtime.GOOS
	}
	// fmt.Println("appMeta.Roaming", appMeta.Roaming)
	// fmt.Println("appMeta.Goos", appMeta.Goos)
	// fmt.Println("appMeta.IsPBaaS", appMeta.IsPBaaS)
	// fmt.Println("appMeta.PBaaS", appMeta.PBaaS)
	if appMeta.Network == "" || appMeta.Network == "." {
		return "."
	}

	if strings.ToLower(appMeta.Network) != "komodo" || strings.ToLower(appMeta.Network) != strings.TrimPrefix(strings.ToLower(appMeta.Network), ".") {
		if strings.ToLower(appMeta.Network) == "vrsctest" {
			if appMeta.IsPBaaS != true {
				appMeta.Network = ".komodo/" + strings.ToLower(appMeta.Network)
			} else {
				appMeta.Network = ".VerusTest/pbaas/" + appMeta.PBaaS
			}
			// appMeta.Network = ".komodo/" + strings.ToLower(appMeta.Network)
		} else {
			appMeta.Network = ".komodo/" + strings.ToUpper(appMeta.Network)
		}
	}

	// The caller really shouldn't prepend the appName with a period, but
	// if they do, handle it gracefully by trimming it.
	// fmt.Println(`strings.TrimPrefix(appMeta.Network, ".")::`, strings.TrimPrefix(appMeta.Network, "."))
	appMeta.Network = strings.TrimPrefix(appMeta.Network, ".")
	appNameUpper := string(unicode.ToUpper(rune(appMeta.Network[0]))) + appMeta.Network[1:]
	appNameLower := string(unicode.ToLower(rune(appMeta.Network[0]))) + appMeta.Network[1:]

	// fmt.Println("appNameLower", appNameLower)
	// fmt.Println("appNameUpper", appNameUpper)

	// Get the OS specific home directory via the Go standard lib.
	var homeDir string
	usr, err := user.Current()
	if err == nil {
		homeDir = usr.HomeDir
	}

	// Fall back to standard HOME environment variable that works
	// for most POSIX OSes if the directory from the Go standard
	// lib failed.
	if err != nil || homeDir == "" {
		homeDir = os.Getenv("HOME")
	}

	switch appMeta.Goos {
	// Attempt to use the LOCALAPPDATA or APPDATA environment variable on
	// Windows.
	case "windows":
		// Windows XP and before didn't have a LOCALAPPDATA, so fallback
		// to regular APPDATA when LOCALAPPDATA is not set.
		appData := os.Getenv("APPDATA")
		if appMeta.Roaming || appData == "" {
			appData = os.Getenv("APPDATA")
		}

		if appData != "" {
			return filepath.Join(appData, appNameUpper)
		}

	case "darwin":
		if homeDir != "" {
			return filepath.Join(homeDir, "Library",
				"Application Support", appNameUpper)
		}

	case "plan9":
		if homeDir != "" {
			return filepath.Join(homeDir, appNameLower)
		}

	default:
		if homeDir != "" {
			return filepath.Join(homeDir, "."+appNameLower)
		}
	}

	// Fall back to the current directory if all else fails.
	return "."
}

// AppDataDir returns an operating system specific directory to be used for
// storing application data for an application.
//
// The appName parameter is the name of the application the data directory is
// being requested for.  This function will prepend a period to the appName for
// POSIX style operating systems since that is standard practice.  An empty
// appName or one with a single dot is treated as requesting the current
// directory so only "." will be returned.  Further, the first character
// of appName will be made lowercase for POSIX style operating systems and
// uppercase for Mac and Windows since that is standard practice.
//
// The roaming parameter only applies to Windows where it specifies the roaming
// application data profile (%APPDATA%) should be used instead of the local one
// (%LOCALAPPDATA%) that is used by default.
//
// Example results:
//  dir := AppDataDir("myapp", false)
//   POSIX (Linux/BSD): ~/.myapp
//   Mac OS: $HOME/Library/Application Support/Myapp
//   Windows: %LOCALAPPDATA%\Myapp
//   Plan 9: $home/myapp
func AppDataDir(appMeta AppMetaData) string {
	return appDataDir(appMeta)
}

// func PBaaSIDHex(appMeta AppMetaData) (string, error) {
// 	// `{"systemtype":"pbaas"}`
// 	// Collect listCurrencies information
// 	_listCurrencies, _ := appMeta.RPCResultMap("listcurrencies", []interface{}{map[string]string{"systemtype": "pbaas"}})
// 	// fmt.Println("_listCurrencies", _listCurrencies)
// 	listCurrencies := _listCurrencies.([]interface{})

// 	// for _, v := range listCurrencies {
// 	// 	if v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["name"] != appMeta.PBaaS {
// 	// 		fmt.Printf("name:id_hex - %v:%v\n", v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["name"], v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["currencyidhex"])
// 	// 		return v.(map[string]interface{})["currencydefinition"].(map[string]interface{})["currencyidhex"].(string), nil
// 	// 	}
// 	// }
// 	return listCurrencies[0].(map[string]interface{})["currencydefinition"].(map[string]interface{})["currencyidhex"].(string), nil

// 	// return "", errors.New("No matching PBaaS found")
// }
