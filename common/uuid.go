package common

import (
	"opsHeart_agent/conf"
	"opsHeart_agent/logger"
	"opsHeart_agent/service/file"
	"opsHeart_agent/utils"
)

var UUID string

// read uuid from /etc/superops_uuid
func InitUUID() {
	uuidF := conf.GetUUIDFile()

	f := file.FDesc{
		Name: uuidF,
	}

	u, err := f.ReadStr()
	if u == "" {
		logger.AgentLog.Infof("action=read uuid;err=%s", err)
		UUID, err = utils.CreateUUID()
		if err != nil {
			logger.AgentLog.Errorf("action=create uuid;err=%s", err)
		}
		_ = f.WriteStr(UUID)
		return
	}
	UUID = u
}
