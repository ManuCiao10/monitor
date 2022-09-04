package main

import (
	"Monitor/browser"
	"Monitor/cfclient"
	"Monitor/constant"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/corpix/uarand"
)

// var Sessions = make(map[string]models.Session)
// var latestVersion = mimic.MustGetLatestVersion(mimic.PlatformWindows)
// var m, _ = mimic.Chromium(mimic.BrandChrome, latestVersion)


func request() {
	start := time.Now()
	
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("Private key cannot be created.", err.Error())
	}

// Generate a pem block with the private key
	keyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	tml := x509.Certificate{
		// you can add any attr that you need
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(5, 0, 0),
		// you have to generate a different serial number each execution
		SerialNumber: big.NewInt(123123),
		Subject: pkix.Name{
			CommonName:   "Awlab",
			Organization: []string{"AwlabCF."},
		},
		BasicConstraintsValid: true,
	}
	// Generate a certificate
	cert, err := x509.CreateCertificate(rand.Reader, &tml, &tml, &key.PublicKey, key)
	if err != nil {
		log.Fatal("Certificate cannot be created.", err.Error())
	}
	// Generate a pem block with the certificate
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})
	if certPem == nil {
		log.Fatal("Certificate cannot be created.")
	}
	// println(string(certPem))
	// println(string(keyPem))

	certicate, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		log.Fatal("Certificate cannot be created.", err.Error())
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{certicate},
			},
		},
	}
	ConfigureClient(client, constant.URL, constant.AGENT)
	req, err := http.NewRequest("GET", constant.URL, nil)
	if err != nil {
		log.Fatal("Request cannot be sent.", err.Error())
	}
	set_headers(req)
	fmt.Print(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(req)
	// log.Println("client: connected to: ", resp.Proto, " server in ", time.Since(start))
	fmt.Printf("<|%v|> [%s]\n", resp.Status, time.Since(start))
	
}

func ConfigureClient(client *http.Client, target string, agent string) error {
	// Initialize the client with the things we need to bypass cloudflare
	cfclient.Initialize(client)

	log.Println("[!] |< Target is protected by Cloudflare, bypassing...|>")

	return browser.GetCloudFlareClearanceCookie(client, agent, target)

}



func set_headers(req *http.Request) {
	req.Header.Set("authority", "en.aw-lab.com")
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", uarand.GetRandom())
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-encoding", "gzip, deflate, br")

}


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