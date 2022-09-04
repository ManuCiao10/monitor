package main

import (
	"encoding/pem"
	"fmt"

	// "io/ioutil"
	"crypto/rand"
	"Monitor/constant"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"log"
	"math/big"
	"net/http"
	"time"
	"github.com/chromedp/chromedp"
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

	certicate, error_cert := tls.X509KeyPair(certPem, keyPem)
	if error_cert != nil {
		log.Fatal("Certificate cannot be created.", error_cert.Error())
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{certicate},
			},
		},
	}
	

	req, err := http.NewRequest("GET", constant.URL, nil)
	if err != nil {
		log.Fatal("Request cannot be sent.", err.Error())
	}
	set_headers(req)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	// fmt.Println(req)
	// log.Println("client: connected to: ", resp.Proto, " server in ", time.Since(start))
	fmt.Printf("<|%v|> [%s]\n", resp.Status, time.Since(start))
	
}


func set_headers(req *http.Request) {
	req.Header.Set("authority", "en.aw-lab.com")
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "dwac_1475e29a8c29e08671dad6a42b=mOUzmsnLCtEUuxd2d-ykMb5NhC2-g63ttps%3D|dw-only|||EUR|false|Europe%2FRome|true; cqcid=abfXR2HZYNGSOE8EJSN8HSGOSN; cquid=||; sid=mOUzmsnLCtEUuxd2d-ykMb5NhC2-g63ttps; dwanonymous_106322550253ae9980e0f038b6061a90=abfXR2HZYNGSOE8EJSN8HSGOSN; __cq_dnt=0; dw_dnt=0; dwsid=4_N1iz6LcTPxu5ek0JzujcVbwjxJw9YPU5v-f02kTMFP9HDiS-dTOz9MvancBBTcfSII97VnioJGlWGGpkkGqA==; _gcl_au=1.1.1131450033.1662174358; _gid=GA1.2.142453763.1662174359; _fphu=%7B%22value%22%3A%225.W1MbhbHTHeJDewJtG8q.1634247081%22%2C%22ts%22%3A1662174359175%7D; __cq_uuid=abYXS961SyxMZaQ91zmrCzFNEw; OptanonAlertBoxClosed=2022-09-03T03:05:59.884Z; _clck=1mrw9b|1|f4k|0; countryMismatch=IT; __cq_bc=%7B%22bclg-awlab-it%22%3A%5B%7B%22id%22%3A%22AW_2212222A%22%2C%22type%22%3A%22vgroup%22%2C%22alt_id%22%3A%22AW_2212222A_8041600%22%7D%5D%7D; __cq_seg=0~-0.20!1~-0.55!2~-0.09!3~0.14!4~0.12!5~-0.14!6~-0.10!7~0.51!8~-0.35!9~-0.45; cto_bundle=vESIOV9ydGFIak5iNDdxUFFoaDcxb0V5OXRqeHBQM0xnTXVVNUlyVlNYJTJGRzF3aVEzSkdNWW5lek1leiUyRlB1OGpBSXlYOEJTMmIwOUpCYjZtRk5xOWZ5aHlYNWJzJTJCZ0FSWnZIU2tCd0VIcGlXUERUSmVVR3JSaTNabWolMkJCeFh0WkklMkJJSzVHdnRLY2N5cnIwbFN2U0VMJTJGNyUyQk5jQlozVzhnVjBibnlmaVVWY0VtVHJBOEI5ZUhLdW84OTk4Q0ZjUFFSa282RA; fanplayr=%7B%22uuid%22%3A%221662174358303-8b31362985c4d7e4d7f9a33a%22%2C%22uk%22%3A%225.W1MbhbHTHeJDewJtG8q.1634247081%22%2C%22sk%22%3A%226d63611035d6e383273bae28a8067c0c%22%2C%22se%22%3A%22e1.fanplayr.com%22%2C%22tm%22%3A1%2C%22t%22%3A1662174571688%7D; __cf_bm=HvdDtItVTlttZnE6ycJZMjD93ZSKG43xXDZqw0RfBZY-1662181422-0-AWBE9zunX21AVJ+cl8+0K3gMvcuLinNpigWSRH0inHgbFJR9VPKYA5mzW6K7umfPKLcMveXnSoT691iIToUTUVc=; _clsk=vzme8v|1662181423240|1|1|h.clarity.ms/collect; cf_chl_2=74baf089309eaee; cf_chl_prog=x13; cf_clearance=FJu4Zeh6etOEe8sw74io19F7YJO0aVmGEJjzBe0QcM8-1662181640-0-150; datadome=NEBZUm-E9zqprNnaMYWJQQFe-2Dt.Kuum7Ju9NSeAY2fl3fPI92eAXRX91mIY5.vDDsZ7TCuOD.k2yW2b4gpsRbS3KIIzdsM2o-cc5thi.mqTUOXDQXdB8Ds05ST33d; OptanonConsent=isGpcEnabled=0&datestamp=Sat+Sep+03+2022+01%3A07%3A26+GMT-0400+(GMT-04%3A00)&version=6.34.0&isIABGlobal=false&hosts=&consentId=a6cdf00a-30fe-4a29-a1d3-448ad80952e2&interactionCount=1&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1%2CC0004%3A1&geolocation=CA%3BQC&AwaitingReconsent=false; _gat_UA-18276494-1=1; _uetsid=5659fec02b3511ed8b88e144a4b88620; _uetvid=1243b4b02d3611ecabd4696673e8ff48; _ga_MVH1E98WW2=GS1.1.1662181646.2.0.1662181646.0.0.0; _ga=GA1.1.1983636562.1662174359")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", constant.AGENT)
	// req.Header.Set("user-agent", uarand.GetRandom())
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-encoding", "gzip, deflate, br")

    
}


func main() {
	request()

}


/*
--- TESTING CF_BYPASS ---
https://nowsecure.nl/
https://gitlab.com/gitlab-com/gl-security/threatmanagement/redteam/redteam-public/cfClearance
https://pkg.go.dev/net/http
*/