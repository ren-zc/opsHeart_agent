package agent

import "testing"

func TestGetAgentInfo(t *testing.T) {
	i, err := GetAgentInfo()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("success: %v", i)
}
