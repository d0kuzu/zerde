package chrome

import (
	"context"
	"github.com/chromedp/chromedp"
	"time"
)

func (c *Client) WaitClick(element string, selector chromedp.QueryOption) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(element, selector),
		chromedp.Click(element, selector),
		chromedp.Sleep(2 * time.Second),
	}
}

func (c *Client) WaitSend(element, text string, selector chromedp.QueryOption) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(element, selector),
		chromedp.SendKeys(element, text, selector),
		chromedp.Sleep(2 * time.Second),
	}
}

func (c *Client) WaitGetText(element string, value *string, selector chromedp.QueryOption) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(element, selector),
		chromedp.Text(element, value, selector),
		chromedp.Sleep(2 * time.Second),
	}
}

func (c *Client) WaitClickRedirect(element string, selector chromedp.QueryOption) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(element, selector),
		chromedp.Click(element, selector),
		chromedp.WaitReady("body"),
	}
}

func (c *Client) getCurrentURL(url *string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		return chromedp.Location(url).Do(ctx)
	})
}

func (c *Client) waitRedirect(initialURL string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		for {
			var cur string
			if err := chromedp.Location(&cur).Do(ctx); err != nil {
				return err
			}
			if cur != initialURL {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		return nil
	})
}
