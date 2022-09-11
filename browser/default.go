package browser

import (
	"Monitor/cfclient"
	"context"
	// "fmt"
	"log"

	"net/http"
	"strings"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"

)

func GetCloudFlareClearanceCookie(client *http.Client, agent string, target string) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// Ignore certificate errors (for use with proxy testing)
		chromedp.Flag("ignore-certificate-errors", "1"),
		// User-Agent MUST match what your tooling uses
		chromedp.UserAgent(agent),
	)
	
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create the chrome instance
	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// Challenges should be solved in ~5 seconds but can be slower. Timeout at 30.
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	// Listen for the Cloudflare cookie
	cookieReceiverChan := make(chan string, 1)
	defer close(cookieReceiverChan)
	// Fetch the login page and wait until CF challenge is solved.
	err := chromedp.Run(ctx,
		chromedp.Navigate(target),
		chromedp.WaitNotPresent(`Checking your browser`, chromedp.BySearch),
		extractCookie(cookieReceiverChan),
	)

	if err != nil {
		return err
	}
	// block the program until the cloud flare cookie is received, or .WaitVisible times out looking for login-pane

	cfToken := <-cookieReceiverChan

	log.Printf("[*] Grabbed Cloudflare token: %s", cfToken)
	
	cookieURL, cookies := cfclient.BakeCookies(target, cfToken)
	client.Jar.SetCookies(cookieURL, cookies)
	return nil
}

func extractCookie(c chan string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		cookies, err := network.GetAllCookies().Do(ctx)
		if err != nil {
			return err
		}
		for _, cookie := range cookies {
			// log.Printf("COOKIE: %#v\n", cookie.Name)
			if strings.Contains(cookie.Name, "__cf_bm") {
				c <- cookie.Value
			} else {
				println("cookies not found\n")
			}
		}
		return nil
	})
}

