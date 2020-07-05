package agent

import (
	"opsHeart_agent/common"
	"opsHeart_agent/service/file"
	"os"
	"runtime"
)

func GetAgentInfo() (*Agent, error) {
	var a Agent
	var err error
	a.Hostname, err = os.Hostname()
	if err != nil {
		return nil, err
	}

	a.UUID = common.UUID

	a.OsArch = runtime.GOARCH
	a.OsType = runtime.GOOS
	a.OsVersion, err = getOsVersion()
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func getOsVersion() (string, error) {
	redHatRealse := "/etc/redhat-release"
	f := file.FDesc{
		Name: redHatRealse,
	}
	return f.ReadStr()
}
