package start_up

import (
	"errors"
	"opsHeart_agent/common"
	"opsHeart_agent/conf"
	"opsHeart_agent/service/agent"
	"opsHeart_agent/service/file"
)

func (rt *RegToken) HandleStatus() error {
	if agent.RegCronStart {
		err := agent.RegCron.Stop()
		if err != nil {
			return err
		}
	}

	if rt.Status == DENIED {
		common.AgentBeDenied = true
		return nil
	}

	if rt.Status == ACCEPTED {
		common.AgentBeDenied = false
		common.SelfToken = rt.Token
		f := file.FDesc{
			Name: conf.GetTokenFile(),
		}
		return f.WriteStr(rt.Token)
	}

	return errors.New("unrecognized status")
}
