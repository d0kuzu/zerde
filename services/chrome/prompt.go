package chrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"html"
	"strings"
	"time"
)

func (c *Client) GetPrompt() (string, error) {
	ctx, cancel := chromedp.NewContext(c.BrowserCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var value string
	tasks := chromedp.Tasks{}
	tasks = append(tasks, chromedp.Navigate("https://diaxel.com/agents/5d5b3880-db37-4b65-a103-db5290e5d5cd/edit"))
	tasks = append(tasks, c.waitGetText(`//*[@id="root"]/div/div/div[2]/main/div[2]/div/div[2]/div[1]/div/div[2]/div/div/div/div/div/div/p`, &value, chromedp.BySearch))
	if err := chromedp.Run(ctx, tasks); err != nil {
		return "", err
	}
	return value, nil
}

func (c *Client) ChangePrompt(prompt string) error {
	ctx, cancel := chromedp.NewContext(c.BrowserCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	prompt = ReplaceNewlinesWithBR(prompt)

	tasks := chromedp.Tasks{}
	tasks = append(tasks, chromedp.Navigate("https://diaxel.com/agents/5d5b3880-db37-4b65-a103-db5290e5d5cd/edit"))
	tasks = append(tasks, chromedp.WaitReady(".ProseMirror"))
	tasks = append(tasks, c.clearProseMirrorText(".ProseMirror"))
	tasks = append(tasks, chromedp.Sleep(500*time.Millisecond))
	tasks = append(tasks, c.setProseMirrorText(".ProseMirror", prompt))
	tasks = append(tasks, chromedp.Sleep(500*time.Millisecond))
	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}

	tasks = chromedp.Tasks{}
	tasks = append(tasks, c.waitClickRedirect(`//*[@id="root"]/div/div/div[2]/main/div[2]/div/div[1]/div/div/div[2]/button`, chromedp.BySearch))
	tasks = append(tasks, chromedp.Sleep(500*time.Millisecond))
	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}
	return nil
}

func ReplaceNewlinesWithBR(s string) string {
	s = html.EscapeString(s)
	s = strings.ReplaceAll(s, "\r\n", "<br>")
	s = strings.ReplaceAll(s, "\r", "<br>")
	s = strings.ReplaceAll(s, "\n", "<br>")

	return s
}
