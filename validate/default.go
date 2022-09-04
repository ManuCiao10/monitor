package validate

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Url(target string) bool {
	u, err := url.Parse(target)

	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func CloudFlareIsPresent(target string, client *http.Client) bool {
	// Check for a typical Cloudflare response
	resp, err := client.Get(target)
	if err != nil {
		log.Fatal("Could not GET target when performing Cloudflare checks")
	}

	if resp.StatusCode == 503 && strings.Contains(resp.Header.Get("Server"), "cloudflare") {
		return true
	}

	return false
}
