package main

import (
	"Monitor/browser"
	"Monitor/cfclient"

	// "Monitor/constant"
	"fmt"

	"errors"
	"log"
	"net/http"

	"gitlab.com/gitlab-com/gl-security/threatmanagement/redteam/redteam-public/cfClearance/validate"
)

func ConfigureClient(client *http.Client, target string, agent string) error {
	// Initialize the client with the things we need to bypass cloudflare
	cfclient.Initialize(client)

	// Validate the target URL
	if validate.Url(target) == false {
		return errors.New("could not parse the target URL")
	}
	fmt.Println("Target URL is valid")
	if validate.CloudFlareIsPresent(target, client) == false {
		log.Println("[*] Target not protected by Cloudflare.")
		return nil
	}

	log.Println("[!] Target is protected by Cloudflare, bypassing...")

	return browser.GetCloudFlareClearanceCookie(client, agent, target)

}

func request() {
	target := "https://www.aw-lab.com"
	userAgent := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"
	myClient := &http.Client{}

	ConfigureClient(myClient, target, userAgent)
	req, _ := http.NewRequest("GET", target, nil)
	req.Header.Set("User-Agent", userAgent)
	resp, err := myClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

}

func main() {
	request()
}

// import (
// 	"C"
// 	"Monitor/constant"
// 	"Monitor/models"
// 	"fmt"
// 	"strings"

// 	// "io/ioutil"
// 	"log"
// 	"net/url"
// 	"time"

// 	// "github.com/corpix/uarand"

// 	// "github.com/PuerkitoBio/goquery"
// 	http "github.com/saucesteals/fhttp"
// 	"github.com/saucesteals/mimic"
// )

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"regexp"
// 	"strconv"

// 	// "html"
// 	"io/ioutil"

// 	"github.com/robertkrimen/otto"
// 	// "github.com/PuerkitoBio/goquery"
// 	// "io/ioutil"
// 	// "io/ioutil"
// )

// // Language: go
// var latestVersion = mimic.MustGetLatestVersion(mimic.PlatformWindows)
// var m, _ = mimic.Chromium(mimic.BrandChrome, latestVersion)

// func initClient(proxy string) (*http.Client, error) {
// 	transport, err := createTransport(proxy)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &http.Client{
// 		Transport: m.ConfigureTransport(transport),
// 	}, nil
// }

// func createTransport(proxy string) (*http.Transport, error) {
// 	if len(proxy) != 0 {
// 		proxyUrl, err := url.Parse(proxy)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &http.Transport{Proxy: http.ProxyURL(proxyUrl)}, nil
// 	} else {
// 		return &http.Transport{}, nil
// 	}
// }

// func request(cParams *C.char) *C.char {
// 	start := time.Now()
// 	var client *http.Client
// 	var req *http.Request

// 	params := C.GoString(cParams)
// 	data := models.SessionParameters{}
// 	err := json.Unmarshal([]byte(params), &data)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	newclient, err := initClient(constant.PROXY)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	client = newclient

// 	req, err = http.NewRequest("GET", constant.URL, nil)
// 	if err != nil {
// 		log.Fatal("Request cannot be sent.", err.Error())
// 	}

// 	transport, err := createTransport(constant.PROXY)
// 	if err != nil {
// 		log.Fatal("Proxy cannot be created.", err.Error())
// 	}
// 	client.Transport = transport

// 	var headerOrder []string
// 	for k, v := range data.Parameters.Headers {
// 		if strings.ToLower(k) != "accept-encoding" && strings.ToLower(k) != "content-length" {
// 			req.Header.Set(k, v)
// 		}
// 		headerOrder = append(headerOrder, k)
// 	}

// 	req.Header[http.HeaderOrderKey] = headerOrder

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	server_header := resp.Header.Get("Server")
// 	if server_header == "cloudflare" && resp.StatusCode == 503 {
// 		log.Println("Cloudflare detected")
// 		n_resp := Solve_cf_challenge(resp, client)
// 		resp = n_resp

// 	}

// 	defer resp.Body.Close()

// 	cookieMap := make(map[string]string)
// 	for _, cookie := range resp.Cookies() {
// 		cookieMap[cookie.Name] = cookie.Value
// 	}
// 	fmt.Printf("<|%v|> [%s]\n", resp.StatusCode, time.Since(start))
// 	// fmt.Printf("<|%v|> \n", resp.Cookies())
// 	fmt.Println(resp)

// 	return C.CString("Finished")
// }

// var jschlRegexp = regexp.MustCompile(`name="jschl_vc" value="(\w+)"`)
// var passRegexp = regexp.MustCompile(`name="pass" value="(.+?)"`)

// func Solve_cf_challenge(resp *http.Response, client *http.Client) *http.Response {
// 	time.Sleep(time.Second * 4)

// 	b, err := ioutil.ReadAll(resp.Body)
// 	resp.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	resp.Body = ioutil.NopCloser(bytes.NewReader(b))

// 	var params = make(url.Values)

// 	if m := jschlRegexp.FindStringSubmatch(string(b)); len(m) > 0 {
// 		params.Set("jschl_vc", m[1])
// 	}

// 	if m := passRegexp.FindStringSubmatch(string(b)); len(m) > 0 {
// 		params.Set("pass", m[1])
// 	}

// 	chkURL, _ := url.Parse("/cdn-cgi/l/chk_jschl")
// 	u := resp.Request.URL.ResolveReference(chkURL)

// 	fmt.Println("Solving challenge for", string(b))
// 	js, err := extractJS(string(b))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	answer, err := evaluateJS(js)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	params.Set("jschl_answer", strconv.Itoa(int(answer)+len(resp.Request.URL.Host)))

// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", u.String(), params.Encode()), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	req.Header.Set("User-Agent", resp.Request.Header.Get("User-Agent"))
// 	req.Header.Set("Referer", resp.Request.URL.String())

// 	log.Printf("Requesting %s?%s", u.String(), params.Encode())
// 	// client := http.Client{
// 	// 	Transport: t.upstream,
// 	// 	Jar:       t.Cookies,
// 	// }

// 	resp, err = client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return resp
// }

// // Language: go
// // Path: go.mod
// var jsRegexp = regexp.MustCompile(
// 	`setTimeout\(function\(\){\s+(var ` +
// 		`s,t,o,p,b,r,e,a,k,i,n,g,f.+?\r?\n[\s\S]+?a\.value =.+?)\r?\n`,
// )
// var jsReplace1Regexp = regexp.MustCompile(`a\.value = (parseInt\(.+?\)).+`)
// var jsReplace2Regexp = regexp.MustCompile(`\s{3,}[a-z](?: = |\.).+`)
// var jsReplace3Regexp = regexp.MustCompile(`[\n\\']`)

// func evaluateJS(js string) (int64, error) {
// 	vm := otto.New()
// 	result, err := vm.Run(js)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return result.ToInteger()
// }

// func extractJS(body string) (string, error) {
// 	matches := jsRegexp.FindStringSubmatch(body)
// 	if len(matches) == 0 {
// 		return "", errors.New("No matching javascript found")
// 	}

// 	js := matches[1]
// 	js = jsReplace1Regexp.ReplaceAllString(js, "$1")
// 	js = jsReplace2Regexp.ReplaceAllString(js, "")

// 	// Strip characters that could be used to exit the string context
// 	// These characters are not currently used in Cloudflare's arithmetic snippet
// 	js = jsReplace3Regexp.ReplaceAllString(js, "")

// 	return js, nil
// }

// func main() {
// 	seshJson := `{"session":"","requestType":"GET","parameters":{"url":"https://www.facebook.com/","proxy":"http://127.0.0.1:8888","headers":{"user-agent":"Go-http-client/2.0","accept-encoding":""},"FORM":null,"JSON":"","cookies":null,"redirects":true},"proxy":""}`
// 	resp := request(C.CString(seshJson))
// 	fmt.Println(C.GoString(resp))
// }

/*
--- TESTING CF_BYPASS ---
https://tls.peet.ws/api/all
https://www.ipify.org/
https://nowsecure.nl/
https://gitlab.com/gitlab-com/gl-security/threatmanagement/redteam/redteam-public/cfClearance
https://pkg.go.dev/net/http
https://privacycheck.sec.lrz.de/
*/
