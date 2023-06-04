package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func MasterSpain(id string, name string, url_input string, image string, country string) {
	fmt.Printf("[%s] Starting thread", id)
	firtRun := true

	var ticketfirst Ticket
	var ticketloop Ticket

	for {
		client := &http.Client{}
		url := fmt.Sprintf("https://availability.ticketmaster.eu/api/v2/TM_ES/availability/%s", id)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Printf("[%s] Error req: %s\n", id, err)
			time.Sleep(5 * time.Second)
		}

		req.Header.Set("authority", "availability.ticketmaster.eu")
		req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8")
		req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
		req.Header.Set("cache-control", "max-age=0")
		req.Header.Set("if-none-match", "66:1:2:1:0")
		req.Header.Set("sec-ch-ua", `"Brave";v="113", "Chromium";v="113", "Not-A.Brand";v="24"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"macOS"`)
		req.Header.Set("sec-fetch-dest", "document")
		req.Header.Set("sec-fetch-mode", "navigate")
		req.Header.Set("sec-fetch-site", "none")
		req.Header.Set("sec-fetch-user", "?1")
		req.Header.Set("sec-gpc", "1")
		req.Header.Set("upgrade-insecure-requests", "1")
		req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
		resp, err := client.Do(req)

		// defer resp.Body.Close()

		if err != nil {
			fmt.Printf("[%s] Error resp: %s\n", id, err)
			time.Sleep(5 * time.Second)
		}

		if resp.StatusCode == 200 {
			bodyText, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("[%s] Error read body: %s\n", id, err)
			}

			if firtRun {
				firtRun = false
				err = json.Unmarshal(bodyText, &ticketfirst)
				if err != nil {
					fmt.Printf("[%s] Error unmarshal: %s\n", id, err)
				}

				fmt.Printf("[%s] [%s] Success got data (firstRun): %v\n", id, resp.Status, ticketfirst)
			} else {
				err = json.Unmarshal(bodyText, &ticketloop)
				if err != nil {
					fmt.Printf("[%s] Error unmarshal: %s\n", id, err)
				}

				fmt.Printf("[%s] [%s] Success got data (loop)\n", id, resp.Status)

				if checkChanged(ticketfirst, ticketloop) {
					fmt.Printf("[%s] Changes detected\n", id)
					ticketfirst = ticketloop

					SendMessage(name, url_input, id, image, ticketfirst)
				} else {
					fmt.Printf("[%s] No changes detected\n", id)
				}

			}

		}

		time.Sleep(10 * time.Second)
	}

}

func SendMessage(name, url, pid, image_url string, ticket Ticket) {
	fmt.Printf("[%s] Sending message\n", pid)
	var fields []Fields
	var array string

	array += fmt.Sprintf("**Region** \n :flag_%s: \n\n", "es")

	if len(ticket.Offers) == 0 {
		array += fmt.Sprintf("**ID: %s\n**PriceLevel: %d\nPriceType: %d\nType: %s\n\n", "PRODUCT OOS", 0, 0, "PRODUCT OOS")
	} else {
		for _, v := range ticket.Offers {
			array += fmt.Sprintf("**ID: %s\n**PriceLevel: %d\nPriceType: %d\nType: %s\n\n", v.ID, v.PriceLevel, v.PriceType, v.Type)
		}
	}

	payload := &Top{
		Username:  "Ticketmaster Spain",
		AvatarURL: img,
		Embeds: []Embeds{
			{
				Title:       name,
				Color:       1999236,
				URL:         url,
				Description: array,
				Fields:      fields,
				Thumbnail: Thumbnail{
					URL: image_url,
				},
				Footer: Footer{
					IconURL: img,
					Text:    "UzumakiToolBox| " + timeWebhook(),
				},
			},
		},
	}

	payloadBuf := new(bytes.Buffer)
	_ = json.NewEncoder(payloadBuf).Encode(payload)

	SendWebhook, err := http.NewRequest("POST", webhook_test, payloadBuf)
	if err != nil {
		fmt.Println(err)
	}
	SendWebhook.Header.Set("content-type", "application/json")

	sendWebhookRes, err := clientWebhook.Do(SendWebhook)
	if err != nil {
		fmt.Print(err)
	}
	if sendWebhookRes.StatusCode != 204 {
		fmt.Printf("Webhook failed with status %d\n", sendWebhookRes.StatusCode)
	}
	defer sendWebhookRes.Body.Close()
}

func timeWebhook() string {
	return time.Now().Format("15:04:05")
}

func checkChanged(first Ticket, loop Ticket) bool {
	if len(first.Offers) != len(loop.Offers) {
		return true
	}

	for i := range first.Offers {
		if first.Offers[i].ID != loop.Offers[i].ID {
			return true
		}
	}

	return false
}
