package hbs

import (
	"opsHeart_agent/logger"
	"opsHeart_agent/utils/cron"
	"time"
)

var Cron *cron.Cr

func init() {
	var err error
	Cron, err = cron.NewCron(RunHbs, nil, 2*time.Minute, 1)
	if err != nil {
		logger.HbsLog.Errorf("acction=init hbs cron;err=%s", err.Error())
	}
}
