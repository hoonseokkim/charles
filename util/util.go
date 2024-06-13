package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"regexp"
	"time"
	// "time"
)

func GetPrivateIp() (string, error) {
	cmd := exec.Command("ifconfig", "en0")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error executing ifconfig command:", err)
		return "", err
	}

	// Define a regular expression to extract the IP address
	re := regexp.MustCompile(`inet\s+(\d+\.\d+\.\d+\.\d+)`)

	// Find the first match
	match := re.FindStringSubmatch(out.String())
	if match == nil {
		fmt.Println("No IP address found")
		return "", err
	}

	// The IP address is the first submatch
	privateIp := match[1]
	fmt.Println("Private IP address:", privateIp)
	return privateIp, nil
}

func Sleep(sec int) {
	sleepDuration := time.Duration(sec) * time.Second
	time.Sleep(sleepDuration)
}