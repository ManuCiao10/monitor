package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"consignment/data"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/patrickmn/go-cache"

	"consignment/webhook"
)

var (
	url  = "https://api-sell.wethenew.com/consignment-slots?skip=0&take=100"
	url1 = "https://api-sell.wethenew.com/consignment-slots?skip=100&take=100"
)

// set headers
func setHeaders(req *http.Request) {
	req.Header.Set("authority", "api-sell.wethenew.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "en-GB,en;q=0.9")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("feature-policy", "microphone 'none'; geolocation 'none'; camera 'none'; payment 'none'; battery 'none'; gyroscope 'none'; accelerometer 'none';")
	req.Header.Set("origin", "https://sell.wethenew.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://sell.wethenew.com/")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Brave";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	req.Header.Set("x-xss-protection", "1;mode=block")
}

func Time() string {
	date := time.Now().Format("15:04:05")
	time := time.Now().UnixNano() / int64(time.Millisecond)
	time_final := fmt.Sprintf("%s.%d", date, time%1000)
	return time_final
}

//create a map string with the url

var urlMap = map[string]string{
	url:  "url",
	url1: "url1",
}

// get products
func getProducts() data.Info {
	var returnInfo data.Info
	var tmpInfo data.Info

	for url := range urlMap {
		options := []tls_client.HttpClientOption{
			tls_client.WithTimeout(30),
			tls_client.WithClientProfile(tls_client.Chrome_105),
			tls_client.WithNotFollowRedirects(),
			// tls_client.WithProxyUrl(discord.GetProxy()),
		}
		client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)

		if err != nil {
			log.Println("Error creating client: ", err)
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		setHeaders(req)
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//append the data to the struct
		err = json.Unmarshal(bodyText, &tmpInfo)
		if err != nil {
			log.Fatal(err)
		}

		returnInfo.Results = append(returnInfo.Results, tmpInfo.Results...)

	}
	return returnInfo
}

// monitor products
func monitorProducts(class data.Info) {
	var new_id data.WTN
	var tmp data.WTN

	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	for _, v := range class.Results {
		c.Set(fmt.Sprintf("%d", v.ID), v.ID, cache.NoExpiration)
	}
	// fmt.Print("Cache: ", c.Items())
	for {
		for url := range urlMap {
			options := []tls_client.HttpClientOption{
				tls_client.WithTimeout(30),
				tls_client.WithClientProfile(tls_client.Chrome_105),
				tls_client.WithNotFollowRedirects(),
				// tls_client.WithProxyUrl(discord.GetProxy()),

			}
			client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
			if err != nil {
				log.Print(err)
			}
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println(err)
			}
			setHeaders(req)

			time.Sleep(3 * time.Second)
			resp, err := client.Do(req)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Printf("[+] Status: <|%d|> %s\n", resp.StatusCode, Time())
			body, _ := io.ReadAll(resp.Body)
			_ = resp.Body.Close()

			if err := json.Unmarshal(body, &tmp); err != nil {
				fmt.Println(err)
			}
			new_id.Results = append(new_id.Results, tmp.Results...)
		}
		// c.Delete(fmt.Sprintf("%d", 596))
		for idx, v := range new_id.Results {
			if _, found := c.Get(fmt.Sprintf("%d", v.ID)); !found {
				c.Set(fmt.Sprintf("%d", v.ID), v.ID, cache.NoExpiration)
				fmt.Println("New Product: ", v.Name)
				webhook.SendWebhook(new_id, idx)
			}
		}
	}

}

func main() {
	products := getProducts()
	monitorProducts(products)

}
