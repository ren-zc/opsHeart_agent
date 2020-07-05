package hbs

import (
	"opsHeart_agent/common"
	"opsHeart_agent/logger"
	"opsHeart_agent/utils/call_http"
	"time"
)

func RunHbs() {
	code, _, err := call_http.HttpGet(common.HbsPath, "")
	if code != 200 || err != nil {
		logger.HbsLog.Errorf("action=send hbs;time=%s;err=%s", time.Now(), err)
	}
}
