package chrome

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"strings"
	"time"
)

func (c *Client) waitClick(element string, selector chromedp.QueryOption) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(element, selector),
		chromedp.Click(element, selector),
		chromedp.Sleep(2 * time.Second),
	}
}

func (c *Client) waitSend(element, text string, selector chromedp.QueryOption) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(element, selector),
		chromedp.SendKeys(element, text, selector),
		chromedp.Sleep(2 * time.Second),
	}
}

func (c *Client) waitGetText(element string, value *string, selector chromedp.QueryOption) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.WaitVisible(element, selector),
		chromedp.Text(element, value, selector),
		chromedp.Sleep(2 * time.Second),
	}
}

func (c *Client) waitClickRedirect(element string, selector chromedp.QueryOption) chromedp.Tasks {
	var url string
	return chromedp.Tasks{
		chromedp.WaitVisible(element, selector),
		c.getCurrentURL(&url),
		chromedp.Click(element, selector),
		c.waitRedirect(&url),
	}
}

func (c *Client) getCurrentURL(url *string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		return chromedp.Location(url).Do(ctx)
	})
}

func (c *Client) waitRedirect(initialURL *string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		for {
			var cur string
			if err := chromedp.Location(&cur).Do(ctx); err != nil {
				return err
			}
			if cur != *initialURL {
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
		return nil
	})
}

func (c *Client) clearProseMirrorText(selector string) chromedp.Action {
	script := fmt.Sprintf(`
	(() => {
		const el = document.querySelector(%q);
		if (!el) return false;

		el.focus();

		const range = document.createRange();
		range.selectNodeContents(el);
		const sel = window.getSelection();
		sel.removeAllRanges();
		sel.addRange(range);
		document.execCommand('delete', false, null);

		return true;
	})()
	`, selector)
	return chromedp.Evaluate(script, nil)
}

func (c *Client) setProseMirrorText(selector, html string) chromedp.Action {
	script := fmt.Sprintf(`
	(() => {
		const el = document.querySelector(%q);
		if (!el) return false;

		el.focus();

		const event = new ClipboardEvent('paste', {
			bubbles: true,
			cancelable: true,
			clipboardData: new DataTransfer()
		});
		event.clipboardData.setData('text/html', %q);
		event.clipboardData.setData('text/plain', %q);
		el.dispatchEvent(event);
		return true;
	})()
	`, selector, html, strings.ReplaceAll(html, "<br>", "\n"))
	return chromedp.Evaluate(script, nil)
}
