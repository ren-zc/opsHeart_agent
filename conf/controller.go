package conf

import "github.com/go-ini/ini"

var cfg *ini.File

func InitCfg() error {
	var err error
	cfg, err = ini.Load("./conf/app.ini")
	return err
}

// GetPort get port config
func GetPort() string {
	return cfg.Section("agent").Key("http_port").String()
}

// GetAddr get addr config
func GetAddr() string {
	return cfg.Section("agent").Key("http_addr").String()
}

func GetLogLevel() string {
	return cfg.Section("app").Key("log_level").String()
}

func GetMode() string {
	return cfg.Section("app").Key("mode").String()
}

func GetServerList() []string {
	return cfg.Section("agent").Key("server_list").Strings(";")
}

func GetUUIDFile() string {
	return cfg.Section("agent").Key("uuid_file").String()
}

func GetTokenFile() string {
	return cfg.Section("agent").Key("token_file").String()
}

func GetSvrUUIDs() []string {
	return cfg.Section("agent").Key("server_uuid").Strings(";")
}

// Used for add or remove server from conf file when server list changed.
func SaveConfData(v string) error {
	cfg.Section("agent").Key("server_list").SetValue(v)
	return cfg.SaveTo("./conf/app.ini")
}
