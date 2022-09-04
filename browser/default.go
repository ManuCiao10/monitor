package browser

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"Monitor/cfclient"
	"log"
	"net/http"
	"strings"
	"time"
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
	println("1")
	// Listen for the Cloudflare cookie
	cookieReceiverChan := make(chan string, 1)
	defer close(cookieReceiverChan)
	println("2")
	// Fetch the login page and wait until CF challenge is solved.
	err := chromedp.Run(ctx,
		chromedp.Navigate(target),
		chromedp.WaitNotPresent(`Checking your browser`, chromedp.BySearch),
		extractCookie(cookieReceiverChan),
	)

	
	println("3")
	if err != nil {
		return err
	}
	println("4")
	// block the program until the cloud flare cookie is received, or .WaitVisible times out looking for login-pane

	cfToken := <-cookieReceiverChan
	
	println("5")
	log.Printf("[*] Grabbed Cloudflare token: %s", cfToken)

	// Finally, build up the cookie jar with the required token
	
	cookieURL, cookies := cfclient.BakeCookies(target, cfToken)
	client.Jar.SetCookies(cookieURL, cookies)
	println("4")
	return nil
}

func extractCookie(c chan string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		cookies, err := network.GetAllCookies().Do(ctx)
		if err != nil {
			return err
		}
		for _, cookie := range cookies {
			if strings.ToLower(cookie.Name) == "cf_clearance" {
				// if we find a proper cookie, put the value on the receiving channel
				c <- cookie.Value
			}
		}
		return nil
	})
}
