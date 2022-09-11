package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

func get(url string, client *http.Client) {
	start := time.Now()
	resp, http_err := client.Get(url)

	if http_err != nil {
		fmt.Printf("Error: %v\n", http_err)
		return
	}

	//Read now to capture time to read full response, not just headers
	_, read_err := ioutil.ReadAll(resp.Body)
	elapsed := time.Since(start)

	if resp != nil {
		defer resp.Body.Close()
	}

	if read_err != nil {
		fmt.Printf("Error: %v, Time taken: %v\n", read_err, elapsed)
		return
	}

	fmt.Printf("Status: %v, Time taken: %v\n", resp.Status, elapsed)
}

func customClient() *http.Client {
	//ref: Copy and modify defaults from https://golang.org/src/net/http/transport.go
	//Note: Clients and Transports should only be created once and reused
	transport := http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			// Modify the time to wait for a connection to establish
			Timeout:   1 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}

	client := http.Client{
		Transport: &transport,
		Timeout:   4 * time.Second,
	}

	return &client
}

func main_2() {
	url := "https://www.aw-lab.com/on/demandware.store/Sites-awlab-it-Site/it_IT/Product-GetAvailability?format=ajax&pid=AW_22121RBA_8041591"
	if len(os.Args) > 1 {
		url = os.Args[1]
	}

	fmt.Println("Fetching: " + url)
	//get(url, http.DefaultClient)
	get(url, customClient())
}