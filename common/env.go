package common

var RegisterPath = "/agent/start-up/register"
var QueryStatus = "/agent/v1/status"
var HbsPath = "/agent/v1/hbs"

var AgentBeDenied = false
var StopRegister chan struct{}

func init() {
	StopRegister = make(chan struct{})
}
