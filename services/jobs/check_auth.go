package jobs

import (
	"AISale/config"
	"fmt"
	"time"
)

func CheckWaitingChats(app *config.App) {
	ticker := time.NewTicker(3 * time.Hour)
	defer ticker.Stop()

	runCheck(app)

	for {
		select {
		case <-app.Chrome.BrowserCtx.Done():
			fmt.Println("JOB Browser context closed — stopping job.")
			return
		case <-ticker.C:
			runCheck(app)
		}
	}
}

func runCheck(app *config.App) {
	if auth, err := app.Chrome.CheckAuth(); err == nil && !auth {
		if err = app.Chrome.UpdateAuth(app.Cfg.DiaxelLogin, app.Cfg.DiaxelPassword); err != nil {
			fmt.Println("JOB UpdateAuth err:", err)
		}
	} else if err != nil {
		fmt.Println("JOB CheckAuth err:", err)
	}
}
