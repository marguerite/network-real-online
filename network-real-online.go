package main

import (
	"fmt"
	"net/http"
	"time"
)

func available() bool {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("Network may be not reachable, waiting...")
		return false
	}
	if resp.StatusCode != 200 {
		fmt.Println("Network seems to be reachable, but is not fully functional yet.")
		return false
	}
	return true
}

func main() {
	ticker := time.NewTicker(time.Second * 1)

	for range ticker.C {
		if available() {
			fmt.Println("Network is reachable, quiting.")
			break
		}
		fmt.Println("Network is still not reachable, wait another second...")
	}
}
