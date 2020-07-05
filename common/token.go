package common

import (
	"opsHeart_agent/conf"
	"opsHeart_agent/logger"
	"opsHeart_agent/service/file"
)

var SelfToken string

func InitToken() {
	tokenF := conf.GetTokenFile()
	f := file.FDesc{
		Name: tokenF,
	}

	var err error
	SelfToken, err = f.ReadStr()
	if err != nil {
		SelfToken = ""
		logger.AgentLog.Errorf("action=read token;err=%s", err)
	}
}
