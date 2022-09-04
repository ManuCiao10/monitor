package cfclient

import (
	"crypto/tls"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func Initialize(client *http.Client) {
	// If a proxy is defined, skip TLS verification.
	// We do this as it seems likely you are testing via ZAP/Burp/etc
	var tr http.Transport
	if os.Getenv("HTTP_PROXY") != "" || os.Getenv("HTTPS_PROXY") != "" {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		tr.Proxy = http.ProxyFromEnvironment
	}

	// Initialize an empty cookie jar. It will be populated later with Cloudflare cookie
	cookieJar, _ := cookiejar.New(nil)

	client.Transport = &tr
	client.Jar = cookieJar
}

func BakeCookies(target string, cfToken string) (*url.URL, []*http.Cookie) {
	u, _ := url.Parse(target)
	d := "." + u.Host
	var cookies []*http.Cookie
	cfCookie := &http.Cookie{
		Name:   "cf_clearance",
		Value:  cfToken,
		Path:   "/",
		Domain: d,
	}
	cookies = append(cookies, cfCookie)
	cookieURL, _ := url.Parse(target)

	return cookieURL, cookies
}
