package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"consignment/data"
)

var (
	webhookURL = "url-webhook"
)

// print time
func timeWebhook() string {
	date := time.Now().Format("15:04:05")
	time := time.Now().UnixNano() / int64(time.Millisecond)
	time_final := fmt.Sprintf("%s.%d", date, time%1000)
	return time_final

}

// sendWebhook
func SendWebhook(new_id data.WTN, idx int) {
	n_size := len(new_id.Results[idx].Sizes)
	var fields []Fields
	for i := 0; i < n_size; i++ {
		fields = append(fields, Fields{
			Name:   "Size",
			Value:  "[" + new_id.Results[idx].Sizes[i] + "](https://sell.wethenew.com/listing/product/" + strconv.Itoa(new_id.Results[idx].ID) + ")",
			Inline: true,
		})
	}
	payload := &Top{
		Username:  "Wethenew Consignment",
		AvatarURL: Image_URL,
		Content:   "New Product To Sell On Wethenew",
		Embeds: []Embeds{
			{
				Title: new_id.Results[idx].Name,
				// Description: "Sell Now",
				Color:  1999236,
				Fields: fields,
				Thumbnail: Thumbnail{
					URL: new_id.Results[idx].Image,
				},
				Footer: Footer{
					IconURL: Image_URL,
					Text:    "Holding-Lab " + timeWebhook(),
				},
			},
		},
	}
	payloadBuf := new(bytes.Buffer)
	_ = json.NewEncoder(payloadBuf).Encode(payload)

	SendWebhook, err := http.NewRequest("POST", webhookURL, payloadBuf)
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
