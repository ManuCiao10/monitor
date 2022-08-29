package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://en.aw-lab.com/on/demandware.store/Sites-awlab-en-Site/en_GB/Product-GetAvailability?format=ajax&pid=AW_106COOCOOA_8012225", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "en.aw-lab.com")
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("cookie", "_gcl_au=1.1.866342104.1661715266; __cq_uuid=abYXS961SyxMZaQ91zmrCzFNEw; OptanonAlertBoxClosed=2022-08-28T19:49:36.959Z; _ga=GA1.2.150548395.1661716177; _gid=GA1.2.1818618316.1661716177; cqcid=adbPEKyPPElIqmCWgqSGaFPhsa; cquid=||; dwanonymous_c705daf0bce3ca74e982743a1044dc05=adbPEKyPPElIqmCWgqSGaFPhsa; __cq_dnt=0; dw_dnt=0; __dfduuid=28c7d337-a59a-47ad-b17e-6264cad911aa; _fbp=fb.1.1661716197473.952419098; grid.alternativeLayout.active=false; _clck=gtmeqs|1|f4f|0; cto_bundle=vtgETF9ydGFIak5iNDdxUFFoaDcxb0V5OXR2RWQzaHNmR1ZzdkxWemxwdmhFUGw0ckdZV3dZbHBLNGFFTmNFV04lMkYlMkY3WUx0V3hUcnVqeWk2VGF1dkVwNyUyQlclMkIzSjBMMEQlMkZhd2dmQSUyRlpnYXdWWlVRZ3NqNFJHQnZFV3JWRUk2Ukc1b3o5ZkpiaWUzRnJuTXdGNlViWGszME05ZyUyQmJmQzhyNDNYRDRzJTJGeUFMcDRvaTVQVlpXc0ZuWXJ3eUpuOEJZWVlvYmFa; cf_clearance=SKmW0FaeH0L0I9tl2yNnYMpUo2ZLS0AOXTmSFqbwG2M-1661808757-0-150; __cf_bm=LP2pVu8DaY1909KhtlkXbylByKhkwfpLY4kHN2jcvek-1661808758-0-ATd0m4Nf4WxQYRKXJTyHrmgH2U7HdgZv3jTxElc4noTCc0R6AZymfYf2QwITDIBPZA3t+QTvA8MCqv+Vj0eW3zQ=; dwac_e14ddc2d119dba9751e7512896=kldQO2TagFgy9wOreCa6xIi-xMpIqSC-v54%3D|dw-only|||EUR|false|Europe%2FRome|true; sid=kldQO2TagFgy9wOreCa6xIi-xMpIqSC-v54; dwsid=3kh6Yx3NX0wd0sM3qjFWwMSh7pIxvv-qkLsM6lwuE8A0jPfappLaXSfBqrOWiL0OLw5oL1ZykC6eQIndp6ze5A==; _fphu=%7B%22value%22%3A%225.W1MbhbHTHeJDewJtG8q.1634247081%22%2C%22ts%22%3A1661808760929%7D; _gat_UA-18276494-1=1; __cq_bc=%7B%22bclg-awlab-en%22%3A%5B%7B%22id%22%3A%22AW_106COOCOOA%22%2C%22type%22%3A%22vgroup%22%2C%22alt_id%22%3A%22AW_106COOCOOA_8012225%22%7D%2C%7B%22id%22%3A%22AW_283DAODAOB%22%2C%22type%22%3A%22vgroup%22%2C%22alt_id%22%3A%22AW_283DAODAOB_5043356%22%7D%2C%7B%22id%22%3A%22AW_22121CZMB%22%2C%22type%22%3A%22vgroup%22%2C%22alt_id%22%3A%22AW_22121CZMB_5017110%22%7D%5D%7D; __cq_seg=0~-0.11!1~-0.42!2~-0.12!3~-0.26!4~0.63!5~0.31!6~-0.21!7~0.19!8~-0.31!9~-0.23!f0~15~5; OptanonConsent=isGpcEnabled=0&datestamp=Mon+Aug+29+2022+17%3A34%3A31+GMT-0400+(GMT-04%3A00)&version=6.34.0&isIABGlobal=false&hosts=&consentId=2a1383ff-b006-4be1-9858-3c1caec520d3&interactionCount=2&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A1%2CC0003%3A1%2CC0004%3A1&geolocation=CA%3BQC&AwaitingReconsent=false; _uetsid=8c081730270a11ed8bb68762a42a05cc; _uetvid=1243b4b02d3611ecabd4696673e8ff48; fanplayr=%7B%22uuid%22%3A%221661715266465-93857204d63ff7ad611ade00%22%2C%22uk%22%3A%225.W1MbhbHTHeJDewJtG8q.1634247081%22%2C%22sk%22%3A%22b2a2e5f0910b294355d189fbf0f036f1%22%2C%22se%22%3A%22e1.fanplayr.com%22%2C%22tm%22%3A1%2C%22t%22%3A1661808871301%7D; _clsk=nqssf0|1661808871748|2|1|h.clarity.ms/collect; datadome=.1XS5C.3VgdB~Y0U.~j3I5tPsoFJi6raM~15SuxKCJaNPUP1Fr~aRh1ZJaiU0GB8evhRAbaK_jTO6jqpRbeFYHYojwOtpZT.fPQiJTn7u3D3HH8g1v9fZhTKx_CO4X41")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://en.aw-lab.com/men/shoes-AW_106COOCOOA.html?dwvar_AW__106COOCOOA_color=8012225")
	req.Header.Set("sec-ch-ua", `"Chromium";v="104", " Not A;Brand";v="99", "Google Chrome";v="104"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
}