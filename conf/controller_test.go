package conf

import "testing"

func TestSetServerList(t *testing.T) {
	_ = InitCfg()
	t.Log(GetServerList())
	err := SaveConfData("0.0.0.0:8080;127.0.0.1:8080")
	if err != nil {
		t.Fatalf("fail: %s\n", err)
	}
	t.Log(GetServerList())
	t.Log(GetSvrUUIDs())
	t.Log("test!")
}
