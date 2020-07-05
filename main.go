package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"opsHeart_agent/common"
	"opsHeart_agent/conf"
	"opsHeart_agent/logger"
	"opsHeart_agent/router"
	"opsHeart_agent/service/agent"
	"opsHeart_agent/service/hbs"
	"opsHeart_agent/utils/call_http"
	"time"
)

func main() {
	// init conf
	err := conf.InitCfg()
	if err != nil {
		fmt.Printf("action=init conf;err=%s\n", err)
		return
	}

	logger.AgentLog.Info("action=load conf file;status=success")

	// init log
	logger.InitLog()

	// init self uuid
	common.InitUUID()

	// init server addr
	common.InitServerAddr()

	// init token
	common.InitToken()

	// check agent if registered
	if common.SelfToken == "" {
		err := agent.RegCron.Start()
		if err != nil {
			logger.AgentLog.Errorf("action=start reg cron;err=%s", err.Error())
			return
		}
		agent.RegCronStart = true
	}

	// if token is not blank, check status
	if common.SelfToken != "" {
		code, b, err := call_http.HttpGet(common.QueryStatus, "")
		if err != nil {
			logger.AgentLog.Errorf("action=check status;err=%s", err.Error())
			return
		}
		if code != 200 {
			err := agent.RegCron.Start()
			if err != nil {
				logger.AgentLog.Errorf("action=check status;do=start reg cron;err=%s", err.Error())
				return
			}
			agent.RegCronStart = true
		} else {
			var r struct {
				Status int `json:"status"`
			}

			err = json.Unmarshal(b, &r)
			if err != nil {
				logger.AgentLog.Errorf("action=check status;do=unmarshal data;err=%s", err.Error())
				return
			}
			switch r.Status {
			case -1:
				logger.AgentLog.Info("action=check status;info=status -1")
				return
			case 2:
				common.AgentBeDenied = false
			default:
				err := agent.RegCron.Start()
				if err != nil {
					logger.AgentLog.Errorf("action=check status;do=start reg cron;err=%s", err.Error())
					return
				}
				agent.RegCronStart = true
			}
		}
	}

	// start register router
	if agent.RegCronStart {
		srv := &http.Server{
			Addr:    conf.GetAddr() + ":" + conf.GetPort(),
			Handler: router.RegRouter,
		}

		go func() {
			logger.AgentLog.Info("action=start register http service;status=success")
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.AgentLog.Errorf("action=start register router;err=%s", err.Error())
				common.StopRegister <- struct{}{}
			}
		}()

		<-common.StopRegister
		err = agent.RegCron.Stop()
		if err != nil {
			logger.AgentLog.Errorf("action=stop register http service;status=false;err=%s", err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			logger.AgentLog.Errorf("action=stop register router;err=%s", err.Error())
			return
		}
		logger.AgentLog.Info("action=stop register http service;status=success")
	}

	if common.AgentBeDenied {
		return
	}

	// start hbs after sleep a random time
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(120)) * time.Second)
	err = hbs.Cron.Start()
	if err != nil {
		logger.HbsLog.Errorf("action=start hbs cron;err=%s", err.Error())
		return
	}
	defer func() {
		_ = hbs.Cron.Stop()
	}()

	logger.AgentLog.Info("action=agent service will be start;")

	// start http server
	err = router.R.Run(fmt.Sprintf("%s:%s", conf.GetAddr(), conf.GetPort()))
	if err != nil {
		logger.AgentLog.Errorf("action=start agent http;err=%s", err.Error())
	}
}
