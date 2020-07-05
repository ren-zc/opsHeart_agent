package file

import "testing"

func TestFDesc_WriteStr(t *testing.T) {
	tmpFile := FDesc{
		Name: "/tmp/echo123.sh",
	}

	err := tmpFile.WriteStr("sleep 3\necho 123\n")
	if err != nil {
		t.Fatalf("failed: %s", err)
	}
}
