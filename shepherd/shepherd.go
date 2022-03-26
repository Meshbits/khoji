package shepherd

import (
	"encoding/json"
	"errors"
	"log"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

func VersionCompare(v1, v2 string) bool {
	upgrade := false

	noV1Str := strings.TrimLeft(v1, "v")
	// fmt.Println("noV1Str:", noV1Str)
	v1Slice := strings.Split(noV1Str, ".")
	// fmt.Println("v1Slice:", v1Slice)
	v1NoDash := strings.Split(v1Slice[2], "-")
	// fmt.Printf("v1NoDash: %v\nv1NoDash Length: %v\n", v1NoDash, len(v1NoDash))

	v1Major, _ := strconv.Atoi(v1Slice[0])
	v1Minor, _ := strconv.Atoi(v1Slice[1])
	v1Patch, _ := strconv.Atoi(v1NoDash[0])
	var v1Build int
	if len(v1NoDash) != 1 {
		v1Build, _ = strconv.Atoi(v1NoDash[1])
	}

	// fmt.Printf("v1Major: %v\nv1Minor: %v\nv1Patch: %v\nv1Build: %v\n", v1Major, v1Minor, v1Patch, v1Build)

	// fmt.Println()

	noV2Str := strings.TrimLeft(v2, "v")
	// fmt.Println("noV2Str:", noV2Str)
	v2Slice := strings.Split(noV2Str, ".")
	// fmt.Println("v2Slice:", v2Slice)
	v2NoDash := strings.Split(v2Slice[2], "-")
	// fmt.Printf("v2NoDash: %v\nv2NoDash Length: %v\n", v2NoDash, len(v2NoDash))

	v2Major, _ := strconv.Atoi(v2Slice[0])
	v2Minor, _ := strconv.Atoi(v2Slice[1])
	v2Patch, _ := strconv.Atoi(v2NoDash[0])
	var v2Build int
	if len(v2NoDash) != 1 {
		v2Build, _ = strconv.Atoi(v2NoDash[1])
	}

	// fmt.Printf("v2Major: %v\nv2Minor: %v\nv2Patch: %v\nv2Build: %v\n\n", v2Major, v2Minor, v2Patch, v2Build)

	// fmt.Println()
	// fmt.Printf("%v < %v: %v\n", v1Minor, v2Minor, v1Minor < v2Minor)
	// fmt.Printf("%v == %v: %v\n", v1Major, v2Major, v1Major == v2Major)
	if v1Major < v2Major {
		// fmt.Println("Major: v2 is new version")
		return upgrade
	}
	if v1Major == v2Major {
		// fmt.Printf("v2 major version (%v) is same as v1 version (%v). Checking minor versions...\n", v2Major, v1Major)
		if v1Minor < v2Minor {
			// fmt.Println("Major: v2 is new version")
			return true
		}
		if v1Minor == v2Minor {
			// fmt.Printf("v2 minor version (%v) is same as v1 version (%v). Checking patch versions...\n", v2Minor, v1Minor)
			if v1Patch < v2Patch {
				// fmt.Println("Patch: v2 is new version")
				return true
			}
			if v1Patch == v2Patch {
				// fmt.Printf("v2 patch version (%v) is same as v1 version (%v). Checking build versions...\n", v2Patch, v1Patch)
				if v1Build < v2Build {
					// fmt.Println("Build: v2 is new version")
					return true
				}
			}
		}
	}
	// fmt.Printf("Upgrade: %v\n", upgrade)
	return upgrade
}

type UpgradeDaemonInfo struct {
	Version        string `json:"version,omitempty"`
	NeedUpdate     bool   `json:"need_update,omitempty"`
	URL            string `json:"url,omitempty"`
	Err            error  `json:"error,omitempty"`
	CurrentVersion string `json:"current_version,omitempty"`
}

func GetDlURL(str ...string) UpgradeDaemonInfo {
	var dlLinuxArm64, dlLinuxAmd64, dlmacOS, dlWin64 string
	// fmt.Println(str)

	client := &fasthttp.Client{}
	url := `https://api.github.com/repos/veruscoin/veruscoin/releases/latest`

	queryByte, err := json.Marshal(map[string]interface{}{
		"id":      "0",
		"jsonrpc": "1.0",
	})
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("GET")
	req.SetBody(queryByte)

	resp := fasthttp.AcquireResponse()
	client.Do(req, resp)

	bodyBytes := resp.Body()
	if len(bodyBytes) != 0 {
		// fmt.Println("bodyBytes len:", len(bodyBytes))
		// fmt.Println("bodyBytes:", string(bodyBytes))

		var res interface{}
		json.Unmarshal(bodyBytes, &res)

		assets := res.(map[string]interface{})["assets"].([]interface{})
		tagName := res.(map[string]interface{})["tag_name"].(string)
		// fmt.Println("tag_name:", tag_name)
		needUpdate := VersionCompare(str[2], tagName)
		if !needUpdate {
			return UpgradeDaemonInfo{Version: tagName, NeedUpdate: true, Err: nil}
		}

		// fmt.Println(`assets -- `, len(assets))
		for i := range assets {
			var arm64 = regexp.MustCompile("(?m)arm64.+$")
			arm64Line := arm64.FindString(assets[i].(map[string]interface{})["name"].(string))
			if arm64Line != "" {
				// fmt.Println(assets[i].(map[string]interface{})["name"])
				// fmt.Println("armLine -", armLine)
				// v1 := strings.TrimLeft(assets[i].(map[string]interface{})["name"].(string), `Verus-CLI-Linux-`)[1:]
				// v2 := strings.TrimRight(v1, `-arm64.tgz`)
				// fmt.Println("version:", v2)
				dlLinuxArm64 = assets[i].(map[string]interface{})["browser_download_url"].(string)
				// fmt.Println("browser_download_url:", assets[i].(map[string]interface{})["browser_download_url"])
			}
			var amd64 = regexp.MustCompile("(?m)x86_64.+$")
			amd64Line := amd64.FindString(assets[i].(map[string]interface{})["name"].(string))
			if amd64Line != "" {
				dlLinuxAmd64 = assets[i].(map[string]interface{})["browser_download_url"].(string)
			}
			var macOS = regexp.MustCompile("(?m)MacOS.+$")
			macOSLine := macOS.FindString(assets[i].(map[string]interface{})["name"].(string))
			if macOSLine != "" {
				dlmacOS = assets[i].(map[string]interface{})["browser_download_url"].(string)
			}
			var win64 = regexp.MustCompile("(?m)Windows.+$")
			win64Line := win64.FindString(assets[i].(map[string]interface{})["name"].(string))
			if win64Line != "" {
				dlWin64 = assets[i].(map[string]interface{})["browser_download_url"].(string)
			}
		}
		// fmt.Println("dlLinuxArm64:", dlLinuxArm64)
		// fmt.Println("dlLinuxAmd64:", dlLinuxAmd64)
		// fmt.Println("dlmacOS:", dlmacOS)
		// fmt.Println("dlWin64:", dlWin64)
		if len(str) != 0 {
			switch str[0] {
			case "darwin":
				return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlmacOS, Err: nil}
			case "linux":
				switch str[1] {
				case "x86_64":
					return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlLinuxAmd64, Err: nil}
				case "amd64":
					return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlLinuxAmd64, Err: nil}
				case "arm64":
					return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlLinuxArm64, Err: nil}
				}
			case "windows":
				return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlWin64, Err: nil}
			}
		} else {
			switch runtime.GOOS {
			case "darwin":
				return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlmacOS, Err: nil}
			case "linux":
				switch runtime.GOARCH {
				case "x86_64":
					return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlLinuxAmd64, Err: nil}
				case "amd64":
					return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlLinuxAmd64, Err: nil}
				case "arm64":
					return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlLinuxArm64, Err: nil}
				}
			case "windows":
				return UpgradeDaemonInfo{Version: tagName, NeedUpdate: needUpdate, URL: dlWin64, Err: nil}
			}
		}
	} else {
		return UpgradeDaemonInfo{Err: errors.New("downloads are unreachable")}
	}
	return UpgradeDaemonInfo{Err: errors.New("something went wrong processing this request")}
}
