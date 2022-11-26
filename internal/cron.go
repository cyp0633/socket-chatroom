package internal

import "github.com/robfig/cron/v3"

func init() {
	c := cron.New()
	c.AddFunc("@every 3s", tryPull)
	c.AddFunc("@every 5s", tryUser)
	c.Start()
}

// 每 5 秒执行一次，从服务器拉取最新状态
func tryPull() {
	if conn != nil {
		DoPull()
	}
}

func tryUser() {
	if conn != nil {
		DoUser()
	}
}
