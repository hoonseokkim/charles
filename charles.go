package main

import (
	"fmt"
	"net/http"
	"net/url"

	"test/charles/constants"
	"test/charles/util"
)

func initChalres() {
	// get private ip
}

func activateThrottling(preset string) {
	cmd, err := makeCharlesCmd(constants.CHARLES_THROTTLING_ACTIVATE, &preset)
	if err != nil {
		return
	}
	
	runCharlesCmd(cmd)
}

func makeCharlesCmd(action string, preset *string) (string, error) {
	cmd := ""
	if preset == nil {
		cmd = constants.CHARLES_CONTROL + "/" + action
	} else {
		cmd = constants.CHARLES_CONTROL + "/" + action + constants.CHARLES_PRESET + *preset
	}
	
	return cmd, nil
}

func runCharlesCmd(cmd string) {
	resp, err := http.Get(cmd)
	if err != nil {
		fmt.Println("[Error] CMD: ", cmd, " error: ", err)
		return
	}

	if resp.StatusCode != 200 {
		fmt.Println("Failed to run CMD: ", cmd)
	}
}

func main() {
	privateIp, err := util.GetPrivateIp()
	if privateIp == "" || err != nil {
		return
	}
	
	chalesUrl := fmt.Sprintf("http://%s:8888", privateIp)
	proxyUrl, err := url.Parse(chalesUrl)
	if err != nil {
		return
	}

	http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	resp, err := http.Get("http://control.charles/throttling/activate?preset=3GG")
	if err != nil {
		fmt.Println("err ", err)
		return
	}

	fmt.Println(resp.StatusCode)
}