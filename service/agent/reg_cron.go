package agent

import (
	"encoding/json"
	"opsHeart_agent/common"
	"opsHeart_agent/logger"
	"opsHeart_agent/utils/call_http"
	"opsHeart_agent/utils/cron"
	"time"
)

var RegCron *cron.Cr
var RegCronStart = false

func RegisterAgent() {
	dat, err := GetAgentInfo()
	if err != nil {
		logger.AgentLog.Errorf("action=register; err=%s", err.Error())
		return
	}

	d, err := json.Marshal(dat)
	if err != nil {
		logger.AgentLog.Errorf("action=register; err=%s", err.Error())
		return
	}
	_, r, err := call_http.HttpPost(common.RegisterPath, d)
	if err != nil {
		logger.AgentLog.Errorf("action=register; err=%s", err.Error())
	} else {
		logger.AgentLog.Infof("action=register; status=success; resp=%s", string(r))
	}
}

func init() {
	var err error
	RegCron, err = cron.NewCron(RegisterAgent, nil, 1*time.Hour, 1)
	if err != nil {
		logger.AgentLog.Errorf("action=init regcron; err=%s", err.Error())
	}
}
