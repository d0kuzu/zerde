package chrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"time"
)

type Client struct {
	AllocCancel context.CancelFunc
	BrowserCtx  context.Context
}

func Init(path string) *Client {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(path),
		chromedp.Flag("headless", false),
		chromedp.Flag("no-sandbox", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.Flag("window-size", "1280,800"),
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	browserCtx, _ := chromedp.NewContext(allocCtx)

	if err := chromedp.Run(browserCtx); err != nil {
		log.Fatalf("Chrome init error: %v", err)
	}

	log.Println("Chrome started (headless)")
	return &Client{BrowserCtx: browserCtx, AllocCancel: allocCancel}
}

func (c *Client) Close() {
	if c.AllocCancel != nil {
		c.AllocCancel()
		log.Println("Waiting Chrome to exit...")
		time.Sleep(500 * time.Millisecond)
	}
}
