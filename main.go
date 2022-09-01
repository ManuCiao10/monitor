package main

import (
	"encoding/pem"
	"fmt"
	"math/big"

	// "io/ioutil"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"time"
	// http "github.com/saucesteals/fhttp"
	// "github.com/saucesteals/mimic"
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
			CommonName:   "New Name",
			Organization: []string{"New Org."},
		},
		BasicConstraintsValid: true,
	}
	cert, err := x509.CreateCertificate(rand.Reader, &tml, &tml, &key.PublicKey, key)
	if err != nil {
		log.Fatal("Certificate cannot be created.", err.Error())
	}
	// Generate a pem block with the certificate
	certPem := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})
	
	tlsCert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
    log.Fatal("Cannot be loaded the certificate.", err.Error())
	}
	l, err := tls.Listen("tcp", ":8080", &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
	})
	


	println(l);
	
	fmt.Println(time.Since(start))


	


	// log.Fatal(s.ListenAndServeTLS("", ""))
	// client := &http.Client{}
	// req, err := http.NewRequest("GET", "https://en.aw-lab.com/on/demandware.store/Sites-awlab-en-Site/en_GB/Product-GetAvailability?format=ajax&pid=AW_106COOCOOA_8012225", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// req.Header.Set("authority", "en.aw-lab.com")
	// req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	// req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
	// req.Header.Set("cache-control", "no-cache")
	// req.Header.Set("content-type", "application/json")
	// req.Header.Set("pragma", "no-cache")
	// req.Header.Set("referer", "https://en.aw-lab.com/men/shoes-AW_106COOCOOA.html?dwvar_AW__106COOCOOA_color=8012225")
	// req.Header.Set("sec-ch-ua", `"Chromium";v="104", " Not A;Brand";v="99", "Google Chrome";v="104"`)
	// req.Header.Set("sec-ch-ua-mobile", "?0")
	// req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	// req.Header.Set("sec-fetch-dest", "empty")
	// req.Header.Set("sec-fetch-mode", "cors")
	// req.Header.Set("sec-fetch-site", "same-origin")
	// req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	// req.Header.Set("x-requested-with", "XMLHttpRequest")
	// resp, err := client.Do(req)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer resp.Body.Close()
	// bodyText, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", bodyText)
	// log.Println("Request took:", time.Since(start))
	// log.Println("Status_code:", resp.StatusCode)
	
}

// func proxies() {
// 	proxies := mimic.GetProxies(mimic.PlatformWindows)
// 	for _, proxy := range proxies {
// 		fmt.Println(proxy)
// 	}
// }

func main() {
	request()

}