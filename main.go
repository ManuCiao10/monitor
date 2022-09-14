package main

import (
	"C"
	"Monitor/constant"
	"Monitor/models"
	"fmt"
	"strings"

	// "io/ioutil"
	"log"
	"net/url"
	"time"

	// "github.com/corpix/uarand"

	// "github.com/PuerkitoBio/goquery"
	http "github.com/saucesteals/fhttp"
	"github.com/saucesteals/mimic"
)
import (
	"bytes"
	"encoding/json"
	// "html"
	"io/ioutil"

	// "github.com/PuerkitoBio/goquery"
	// "io/ioutil"
	// "io/ioutil"
)

// Language: go
var latestVersion = mimic.MustGetLatestVersion(mimic.PlatformWindows)
var m, _ = mimic.Chromium(mimic.BrandChrome, latestVersion)

func request(cParams *C.char) *C.char {
	start := time.Now()
	var client *http.Client
	var req *http.Request

	params := C.GoString(cParams)
	data := models.SessionParameters{}
	err := json.Unmarshal([]byte(params), &data)
	if err != nil {
		log.Println(err)
	}

	newclient, err := initClient(constant.PROXY)
	if err != nil {
		log.Fatal(err)
	}
	client = newclient

	req, err = http.NewRequest("GET", constant.URL, nil)
	if err != nil {
		log.Fatal("Request cannot be sent.", err.Error())
	}

	transport, err := createTransport(constant.PROXY)
	if err != nil {
		log.Fatal("Proxy cannot be created.", err.Error())
	}
	client.Transport = transport

	var headerOrder []string
	for k, v := range data.Parameters.Headers {
		if strings.ToLower(k) != "accept-encoding" && strings.ToLower(k) != "content-length" {
			req.Header.Set(k, v)
		}
		headerOrder = append(headerOrder, k)
	}

	req.Header[http.HeaderOrderKey] = headerOrder

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	server_header := resp.Header.Get("Server")
	if server_header == "cloudflare" && resp.StatusCode == 503 {
		log.Println("Cloudflare detected")
		n_resp := Solve_cf_challenge(resp)
		resp = n_resp

	}

	defer resp.Body.Close()

	cookieMap := make(map[string]string)
	for _, cookie := range resp.Cookies() {
		cookieMap[cookie.Name] = cookie.Value
	}
	fmt.Printf("<|%v|> [%s]\n", resp.StatusCode, time.Since(start))
	// fmt.Printf("<|%v|> \n", resp.Cookies())

	return C.CString("Finished")
}

func Solve_cf_challenge(resp *http.Response) *http.Response {
	time.Sleep(time.Second * 4)

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	resp.Body = ioutil.NopCloser(bytes.NewReader(b))

	return resp
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

// func ConfigureClient(client *http.Client, target string, agent string) error {
// 	// Initialize the client with the things we need to bypass cloudflare
// 	cfclient.Initialize(client)

// 	log.Println("[!] |< Target is protected by Cloudflare, bypassing...|>")

// 	return browser.GetCloudFlareClearanceCookie(client, agent, target)

// }

// }

func main() {
	seshJson := `{"session":"","requestType":"GET","parameters":{"url":"https://www.facebook.com/","proxy":"http://127.0.0.1:8888","headers":{"user-agent":"Go-http-client/2.0","accept-encoding":""},"FORM":null,"JSON":"","cookies":null,"redirects":true},"proxy":""}`
	resp := request(C.CString(seshJson))
	fmt.Println(C.GoString(resp))
}

/*
--- TESTING CF_BYPASS ---
https://tls.peet.ws/api/all
https://www.ipify.org/
https://nowsecure.nl/
https://gitlab.com/gitlab-com/gl-security/threatmanagement/redteam/redteam-public/cfClearance
https://pkg.go.dev/net/http
https://privacycheck.sec.lrz.de/
*/
