package common

import (
	"math/rand"
	"opsHeart_agent/conf"
	"opsHeart_agent/logger"
	"time"
)

var (
	UsedServer string
	addrList   []string
)

func InitServerAddr() {
	addrList = conf.GetServerList()

	l := len(addrList)

	if l == 0 {
		UsedServer = ""
		logger.AgentLog.Warningf("action=init server addr; err=no server found in config")
	}

	if l == 1 {
		UsedServer = addrList[0]
	}

	if l > 1 {
		rand.Seed(time.Now().Unix())
		UsedServer = addrList[rand.Intn(l)]
	}
}
