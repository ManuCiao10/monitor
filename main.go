package main

import (
	"Monitor/browser"
	"Monitor/cfclient"
	"Monitor/constant"
	"Monitor/models"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"log"
	// "math/big"
	// "net/http"
	// "net/url"
	"time"

	http "github.com/saucesteals/fhttp"
	"github.com/saucesteals/mimic"
)

// Language: go
var latestVersion = mimic.MustGetLatestVersion(mimic.PlatformWindows)
var m, _ = mimic.Chromium(mimic.BrandChrome, latestVersion)

func request() {
	start := time.Now()
	// Create a new client
	var client *http.Client
	// Create a new request
	var req *http.Request

	client , err := initClient(constant.PROXY)
	if err != nil {
		log.Fatal(err)
	}



	// req, err := http.NewRequest("GET", constant.URL, nil)
	// if err != nil {
	// 	log.Fatal("Request cannot be sent.", err.Error())
	// }

	// transport, err := createTransport(constant.PROXY)
	// if err != nil {
	// 	log.Fatal("Proxy cannot be created.", err.Error())
	// }
	// client.Transport = transport

	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(body))
	fmt.Printf("<|%v|> [%s]\n", resp.Status, time.Since(start))

}


func initClient(proxy string) (*http.Client, error) {
	transport, err := createTransport(proxy)
	if err != nil {
		return nil, err
	}

	return &http.Client{
		Transport: m.ConfigureTransport(transport),
	}, nil
}



func createTransport(proxy string) (*http.Transport, error) {
	if len(proxy) != 0 {
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}
		return &http.Transport{Proxy: http.ProxyURL(proxyUrl)}, nil
	} else {
		return &http.Transport{}, nil
	}
}

// func createTransport(proxy string) (*http.Transport, error) {
// 	transport := &http.Transport{
// 		TLSClientConfig: &tls.Config{
// 			InsecureSkipVerify: true,
// 		},
// 	}
// 	if proxy != "" {
// 		proxyURL, err := url.Parse(proxy)
// 		if err != nil {
// 			return nil, err
// 		}
// 		transport.Proxy = http.ProxyURL(proxyURL)
// 	}
// 	return transport, nil
// }

func ConfigureClient(client *http.Client, target string, agent string) error {
	// Initialize the client with the things we need to bypass cloudflare
	cfclient.Initialize(client)

	log.Println("[!] |< Target is protected by Cloudflare, bypassing...|>")

	return browser.GetCloudFlareClearanceCookie(client, agent, target)

}

// func set_headers(req *http.Request) {
// 	req.Header.Set("authority", "en.aw-lab.com")
// 	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
// 	req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
// 	req.Header.Set("cache-control", "no-cache")
// 	req.Header.Set("content-type", "application/json")
// 	req.Header.Set("pragma", "no-cache")
// 	req.Header.Set("sec-fetch-dest", "empty")
// 	req.Header.Set("sec-fetch-mode", "cors")
// 	req.Header.Set("sec-fetch-site", "same-origin")
// 	req.Header.Set("user-agent", uarand.GetRandom())
// 	req.Header.Set("x-requested-with", "XMLHttpRequest")
// 	req.Header.Set("accept", "*/*")
// 	req.Header.Set("accept-encoding", "gzip, deflate, br")

// }

// fmt.Println(req)
// fmt.Println(resp)
// cookieMap := make(map[string]string)
// for _, cookie := range resp.Cookies() {
// 	cookieMap[cookie.Name] = cookie.Value
// }
// fmt.Println(cookieMap)
// log.Println("client: connected to: ", resp.Proto, " server in ", time.Since(start))

func main() {
	request()
}

/*
--- TESTING CF_BYPASS ---
https://tls.peet.ws/api/all
https://nowsecure.nl/
https://gitlab.com/gitlab-com/gl-security/threatmanagement/redteam/redteam-public/cfClearance
https://pkg.go.dev/net/http
*/
