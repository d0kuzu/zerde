package chrome

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/chromedp"
	"time"
)

func (c *Client) CheckAuth() (bool, error) {
	ctx, cancel := chromedp.NewContext(c.BrowserCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	originUrl := "https://diaxel.com/login"
	var url string
	tasks := chromedp.Tasks{}
	tasks = append(tasks, chromedp.Navigate(originUrl))
	tasks = append(tasks, chromedp.Sleep(10*time.Second))
	tasks = append(tasks, chromedp.Location(&url))

	if err := chromedp.Run(ctx, tasks); err != nil {
		return false, errors.New("JOB CheckAuth err: " + err.Error())
	}
	return originUrl != url, nil
}

func (c *Client) UpdateAuth(login, password string) error {
	ctx, cancel := chromedp.NewContext(c.BrowserCtx)
	defer cancel()
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	tasks := chromedp.Tasks{}
	tasks = append(tasks, chromedp.Navigate("https://diaxel.com/login"))
	tasks = append(tasks, c.waitSend(`email`, login, chromedp.ByID))
	tasks = append(tasks, c.waitSend("password", password, chromedp.ByID))
	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}
	fmt.Println("Passed step 1")

	tasks = chromedp.Tasks{}
	tasks = append(tasks, c.waitClickRedirect(`//*[@id="root"]/div/div/div/div[1]/div/form/div[4]/button`, chromedp.BySearch))
	tasks = append(tasks, chromedp.Sleep(500*time.Millisecond))
	if err := chromedp.Run(ctx, tasks); err != nil {
		return err
	}
	fmt.Println("Passed step 2")
	return nil
}
