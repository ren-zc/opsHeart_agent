package agent

type Agent struct {
	UUID         string `json:"uuid"`
	Hostname     string `json:"hostname"`
	OsType       string `json:"os_type"`
	OsVersion    string `json:"os_version"`
	OsArch       string `json:"os_arch"`
	AgentVersion string `json:"agent_version"`
}
