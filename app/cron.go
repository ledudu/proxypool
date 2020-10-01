package app

import (
	"github.com/jasonlvhit/gocron"
)

func Cron() {
	_ = gocron.Every(120).Minutes().Do(CrawlGo)
	<-gocron.Start()
}
